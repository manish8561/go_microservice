import BaseModel from "../../model/base.model"
import { blockDiff, chainIdArr, eventsArr, initial } from "../../bin/token"
import ProposalCreated from "./schema/ProposalCreated"
import governanceAbi from "../../bin/governance.abi.json"
import * as Helpers from '../../helpers';


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
        console.log(count, event, "==================");
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
      for (let chainId of chainIdArr) {
        await this.getCronLogs(chainId);
      }
    } catch (error) {
      console.log(error, '============================================');
    }
  }

  /**
   * delete the duplicates
   * @param  {string} d
   */
  private async callingDelete(d: string, obModel: any) {
    let query: any = [];

    switch (d) {
      // case "ManagementPoolRegistration":
      //   query = [
      //     {
      //       $group: {
      //         _id: {
      //           "chainId":"$chainId",
      //           transactionHash: "$transactionHash",
      //           user: "$user",
      //           poolType: "$poolType",
      //         },
      //         dups: { $addToSet: "$_id" },
      //         count: { $sum: 1 },
      //       },
      //     },
      //     { $match: { count: { $gt: 1 } } },
      //   ];
      //   break;

      case "ProposalCreated":
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
   * get cronlogs
   * @param  {Number} chainId
   */
  private async getCronLogs(chainId: Number) {
    try {
      //getting the current block number
      const currentBlockNumber = await Helpers.Web3Helper.getBlockNumber(chainId);
      for (let event of eventsArr) {
        const ob = (await import(`./schema/${event}`)).default;
        const row: any = await this.getLastRecord(ob, chainId);
        if (row) {
          let abi = governanceAbi;
          // if (row.contractName === "governance") {
          //   abi = governanceAbi;
          // }
          const contractObj = await Helpers.Web3Helper.callContract(row.chainId, abi, row.contract);
          const d = {
            contractObj,
            event,
            fromBlock: row.lastBlockNumber
          };
          console.log(d.fromBlock, d.event, 'before event')
          let eventData = await Helpers.Web3Helper.getEvents(d);
          let filterData: any = [];

          if (eventData && eventData.length === 0) {
            let l = row.lastBlockNumber + blockDiff;
            if (l >= currentBlockNumber) {
              l = currentBlockNumber;
            }
            row.lastBlockNumber = l;
            await row.save();
            continue;
          } else {
            if (eventData.length === 1 && row.blockNumber === eventData[0].blockNumber) {
              let l = row.lastBlockNumber + blockDiff;
              if (l >= currentBlockNumber) {
                l = currentBlockNumber;
              }
              row.lastBlockNumber = l;
              await row.save();
              continue;
            }
            console.log('hi')
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
    } catch (error) {
      throw error;
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
        returnValues.id = Number(obj.returnValues.id);
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
}

export default new EventModel();