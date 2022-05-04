import mongoose, { Schema } from 'mongoose';
class ProposalCreatedSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }
    /* 
     // @notice An event emitted when a new proposal is created
   
    // @notice An event emitted when a new proposal is created
        event ProposalCreated(
            uint256 id,
            address proposer,
            address[] targets,
            uint256[] values,
            string[] signatures,
            bytes[] calldatas,
            uint256 startTime,
            uint256 endTime,
            string description,
            uint256 proposalType
        );

        // @notice An event emitted when the first vote is cast in a proposal
        event StartBlockSet(uint256 proposalId, uint256 startBlock);

        // @notice An event emitted when a vote has been cast on a proposal
        event VoteCast(
            address voter,
            uint256 proposalId,
            bool support,
            uint256 votes
        );

        // @notice An event emitted when a proposal has been canceled
        event ProposalCanceled(uint256 id);

        // @notice An event emitted when a proposal has been queued in the Timelock
        event ProposalQueued(uint256 id, uint256 eta);

        // @notice An event emitted when a proposal has been executed in the Timelock
        event ProposalExecuted(uint256 id);

    */

    private schema() {
        this.objectSchema = new Schema({
            chainId: { type: Number },
            transactionHash: { type: String },
            blockNumber: { type: Number, default: 0 },
            lastBlockNumber: { type: Number, default: 0 },
            contractName: { type: String, required: true },
            contract: { type: String, required: true },
            id: { type: Number },// proposal id in contract
            proposer: { type: String },
            targets: { type: [String], default: [] },
            values: { type: [String], default: [] },
            signatures: { type: [String], default: [] },
            calldatas: { type: [String], default: [] },
            startTime: { type: Number },
            endTime: { type: Number },
            description: { type: String },
            proposalType: { type: Number },
            cronStatus: { type: String, default: "pending" }

        }, { timestamps: false, strict: false });

        this.objectSchema.index({ contract: 1, chainId: 1, blockNumber: -1, });
    }
}

export default mongoose.model('ProposalCreated', (new ProposalCreatedSchema()).objectSchema);