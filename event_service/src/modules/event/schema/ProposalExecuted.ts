import mongoose, { Schema } from 'mongoose';
class ProposalExecutedSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }
    /* 
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

            proposalId: { type: Number },
        }, { timestamps: false, strict: false });

        this.objectSchema.index({ contract: 1, chainId: 1, blockNumber: -1, });
    }
}

export default mongoose.model('ProposalExecuted', (new ProposalExecutedSchema()).objectSchema);