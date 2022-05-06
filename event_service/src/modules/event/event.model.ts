import BaseModel from "../../model/base.model"
import { blockDiff, chainIdArr, eventsArr, initial } from "../../bin/token"
import ProposalCreated from "./schema/ProposalCreated"
import governanceAbi from "../../bin/governance.abi.json"
import * as Helpers from '../../helpers'
import Proposal from "./schema/Proposal"

class EventModel extends BaseModel {

  constructor() {
    super();
  }
  /**
   * checking the add of event
   */
  public async check() {
    try {
      const ob: any = new ProposalCreated();

      ob.cronStatus = "pending";
      ob.transactionHash = "transaction hash";
      ob.blockNumber = 123434;
      ob.lastBlockNumber = 123434;
      ob.id = 1;
      ob.governance = "governance address";
      ob.proposer = "propser address";
      ob.targets = ["test", "test2"];
      ob.values = ["test", "test2"];
      ob.signatures = ["test", "test2"];
      ob.calldatas = ["test", "test2"];
      ob.startTime = (new Date()).getTime();
      ob.endTime = (new Date()).getTime();
      ob.description = "ipfs hash";
      ob.proposalType = 2;
      return await ob.save();
    } catch (error) {
      throw error;
    }
  }
  /**
   * intialize collection
   */
  public async initialize() {
    for (let event of eventsArr) {
      try {
        const ob = (await import(`./schema/${event}`)).default;
        const count = await ob.countDocuments({});
        // console.log(count, event, "==================");
        if (count === 0) {
          console.log("zero")
          const initialData = initial[event];
          if (initialData && initialData.length > 0) {
            await this.addData(ob, initialData);
          }
        }
      } catch (error) {
        console.log("Error while initialize the events", error, "================");
      }
    }
  }
  /**
   * add Data
   * @param  {any} obModel
   * @param  {any} data
   * @returns Promise
   */
  private async addData(obModel: any, data: any): Promise<any> {
    try {
      return await obModel.insertMany(data)
    } catch (error) {
      throw error;
    }
  }
  /**
   * getLastRecord
   * @param  {any} obModel
   * @returns Promise
   */
  private async getLastRecord(obModel: any, chainId: Number): Promise<any> {
    try {
      return await obModel.findOne({ chainId }).sort({ blockNumber: -1 });
    } catch (error) {
      throw error;
    }
  }

  /**
   * get logs
   */
  public async getLogs() {
    try {
      let a: any = [];
      for (let chainId of chainIdArr) {
        a = [...a, this.getCronLogs(chainId)];
      }
      //calling function in parallel with chain id
      await Promise.all(a);
    } catch (error) {
      console.log(error, '============================================');
    }
  }
  /**
   * get cronlogs
   * @param  {Number} chainId
   */
  private async getCronLogs(chainId: Number) {
    try {
      //getting the current block number
      const currentBlockNumber = await Helpers.Web3Helper.getBlockNumber(chainId);
      let start = Date.now();
      let a: any = [];
      for (let event of eventsArr) {
        a = [...a, this.getEventLogs(event, chainId, currentBlockNumber)]
      }
      // parallel call for each event
      await Promise.all(a);
      console.log((Date.now() - start) / 1000, "after promise all: ");

    } catch (error) {
      throw error;
    }
  }
  /**
   * @param  {string} event
   * @param  {Number} chainId
   * @returns Promise
   */
  private async getEventLogs(event: string, chainId: Number, currentBlockNumber: Number): Promise<void> {
    const ob = (await import(`./schema/${event}`)).default;
    const row: any = await this.getLastRecord(ob, chainId);
    if (row) {
      let abi = governanceAbi;
      if (row.contractName === "governance") {
        abi = governanceAbi;
      }
      const contractObj = await Helpers.Web3Helper.callContract(row.chainId, abi, row.contract);
      const d = {
        contractObj,
        event,
        fromBlock: row.lastBlockNumber
      };
      let eventData = await Helpers.Web3Helper.getEvents(d);
      let filterData: any = [];

      if (eventData && eventData.length === 0) {
        let l = row.lastBlockNumber + blockDiff;
        if (l >= currentBlockNumber) {
          l = currentBlockNumber;
        }
        row.lastBlockNumber = l;
        await row.save();

        return;
      } else {
        if (eventData.length === 1 && row.blockNumber === eventData[0].blockNumber) {
          let l = row.lastBlockNumber + blockDiff;
          if (l >= currentBlockNumber) {
            l = currentBlockNumber;
          }
          row.lastBlockNumber = l;
          await row.save();
          return;
        }
        if (eventData.length > 0 && row.blockNumber === eventData[0].blockNumber) {
          eventData = eventData.slice(1);
        }
        //filterTheArray
        for (let d of eventData) {
          d = await this.filterObjectValues(d, row, event);
          filterData = [...filterData, d];
        }
        //adding the data
        let r = await this.addData(ob, filterData);
        r = await this.callingDelete(event, ob);
        console.log("inserted", r)
      }
      // console.log(eventData, filterData, "in loop")
    }
  }
  /**
   * delete the duplicates
   * @param  {string} d
   */
  private async callingDelete(d: string, obModel: any) {
    let query: any = [];

    switch (d) {
      case "ProposalCreated":
      case 'ProposalCanceled':
      case 'ProposalExecuted':
      case 'ProposalQueued':
      case 'StartBlockSet':
      case 'VoteCast':

        query = [
          {
            $group: {
              _id: {
                chainId: "$chainId",
                transactionHash: "$transactionHash",
              },
              dups: { $addToSet: "$_id" },
              count: { $sum: 1 },
            },
          },
          { $match: { count: { $gt: 1 } } },
        ];
        break;
    }

    if (query.length > 0) {
      try {
        const result = await obModel
          .aggregate(query);
        if (result.length > 0) {
          result.forEach(async (doc: any) => {
            doc.dups.shift();
            await obModel.deleteMany({ _id: { $in: doc.dups } });
          });
        }
      } catch (error) {
        throw error;
      }
    }
  }

  /**
   * filter array object properties
   * @param  {any} obj
   * @param  {string} eventName
   */
  private async filterObjectValues(obj: any, row: any, eventName: string) {
    const returnValues: any = {};

    returnValues.chainId = row.chainId;
    returnValues.contractName = row.contractName;
    returnValues.contract = row.contract;
    returnValues.transactionHash = obj.transactionHash;
    returnValues.blockNumber = obj.blockNumber;
    returnValues.lastBlockNumber = obj.blockNumber;

    switch (eventName) {
      case 'ProposalCreated':
        returnValues.proposalId = Number(obj.returnValues.id);
        returnValues.proposer = obj.returnValues.proposer.toLowerCase();
        returnValues.targets = obj.returnValues.targets;
        returnValues.values = obj.returnValues.values;
        returnValues.signatures = obj.returnValues.signatures;
        returnValues.calldatas = obj.returnValues.calldatas;
        returnValues.startTime = Number(obj.returnValues.startTime);
        returnValues.endTime = Number(obj.returnValues.endTime);
        returnValues.description = obj.returnValues.description;
        returnValues.proposalType = Number(obj.returnValues.proposalType);
        break;
      case 'ProposalCanceled':
        returnValues.proposalId = Number(obj.returnValues.proposalId);
        break;
      case 'ProposalExecuted':
        returnValues.proposalId = Number(obj.returnValues.id);
        break;
      case 'ProposalQueued':
        returnValues.proposalId = Number(obj.returnValues.id);
        returnValues.eta = Number(obj.returnValues.eta);
        break;
      case 'StartBlockSet':
        returnValues.proposalId = Number(obj.returnValues.proposalId);
        returnValues.startBlock = Number(obj.returnValues.startBlock);
        break;
      case 'VoteCast':
        returnValues.proposalId = Number(obj.returnValues.proposalId);
        returnValues.voter = obj.returnValues.voter.toLowerCase();
        returnValues.support = obj.returnValues.support;
        returnValues.votes = Number(obj.returnValues.votes / 10 ** 18);
        break;
      // case 'InvestmentRewardLog':
      //   returnValues.user = obj.returnValues.user.toLowerCase();
      //   returnValues.time = Number(obj.returnValues.time);
      //   returnValues.investmentAmount = Number(obj.returnValues.investmentAmount / 10 ** 4);
      //   returnValues.investmentToken = Number(obj.returnValues.investmentToken / 10 ** 22);
      //   returnValues.planType = obj.returnValues.planType;
      //   break;
    }
    return returnValues;
  }
  /**
   * updateProposalStatus
   * @param  {Number} proposalId
   * @returns Promise
   */
  public async updateProposalStatus(docs: any, event: string): Promise<void> {
    // handle doc inserted in the event collections
    switch (event) {
      case "ProposalCreated":
        for (let d of docs) {
          const record: any = await Proposal.findOne({ chain_id: d.chainId, transaction_hash: d.transactionHash, proposal_id: d.proposalId });
          if (record) {
            // update the proposal if found
            record.start_time = d.startTime;
            record.end_time = d.endTime;
            record.proposer = d.proposer;
            record.description = d.decription;
            record.proposal_type = d.proposalType;
            await record.save()

          } else {
            //insert proposal if not found.
            const record: any = new Proposal();
            record.chain_id = d.chainId;
            record.transaction_hash = d.transactionHash;
            record.block_number = d.blockNumber;
            record.proposal_id = d.proposalId;
            record.start_time = d.startTime;
            record.end_time = d.endTime;
            record.proposer = d.proposer;
            record.description = d.description;
            record.proposal_type = d.proposalType;
            record.voting_period = 0;
            record.eta = 0;
            record.for_votes = 0;
            record.against_votes = 0;
            record.canceled = false;
            record.executed = false;
            record.title = '';
            record.db_description = '';
            record.status = "Pending";
            record._created = new Date(),
              record._modified = new Date(),

              await record.save();
          }
        }
        break;
      case "StartBlockSet":
      case "ProposalCanceled":
      case "ProposalExecuted":
        for (let d of docs) {
          const record: any = await Proposal.findOne({ chain_id: d.chainId, proposal_id: d.proposalId });
          if (record) {
            const status = await this.getProposalState(d);
            // update the proposal if found
            record.status = status;
            await record.save();
          }
        }
        break;
      case "ProposalQueued":
        for (let d of docs) {
          const record: any = await Proposal.findOne({ chain_id: d.chainId, proposal_id: d.proposalId });
          if (record) {
            const status = await this.getProposalState(d);
            // update the proposal if found
            record.status = "Queued";
            record.eta = d.eta;
            await record.save();
          }
        }
        break;
      case "VoteCast":
        for (let d of docs) {
          const record: any = await Proposal.findOne({ chain_id: d.chainId, proposal_id: d.proposalId });
          if (record) {
            const p = await this.getProposalContract(d);
            if (p) {
              // update the proposal if found
              record.for_votes = (p.forVotes) / (10 ** 18);
              record.against_votes = (p.againstVotes) / (10 ** 18);
              await record.save();
            }
          }
        }
        break;
    }
  }
  /**
   * state proposal
   * @param  {any} doc
   * @returns Promise
   *    Pending,
        Active,
        Canceled,
        Defeated,
        Succeeded,
        Queued,
        Expired,
        Executed

   */
  public async getProposalState(doc: any): Promise<string> {
    let abi = governanceAbi;
    if (doc.contractName === "governance") {
      abi = governanceAbi;
    }
    const contractObj: any = await Helpers.Web3Helper.callContract(doc.chainId, abi, doc.contract);
    const r = await contractObj.methods.state(doc.proposalId).call();
    // console.log(r, 'state')
    let state = 'Pending';
    switch (Number(r)) {
      case 0:
        state = "Pending";
        break;
      case 1:
        state = "Active";
        break;
      case 2:
        state = "Canceled";
        break;
      case 3:
        state = "Defeated";
        break;
      case 4:
        state = "Succeeded";
        break;
      case 5:
        state = "Queued";
        break;
      case 6:
        state = "Expired";
        break;
      case 7:
        state = "Executed";
        break;
    }

    return state;
  }
  /**
   * get proposal struct
   * @param  {any} doc
   * @returns Promise
   */
  public async getProposalContract(doc: any): Promise<any> {
    let abi = governanceAbi;
    if (doc.contractName === "governance") {
      abi = governanceAbi;
    }
    try {
      const contractObj: any = await Helpers.Web3Helper.callContract(doc.chainId, abi, doc.contract);
      return await contractObj.methods.proposals(doc.proposalId).call();
    } catch (error) {
      throw error;
    }

  }
}

export default new EventModel();