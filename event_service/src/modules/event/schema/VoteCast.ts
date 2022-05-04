import mongoose, { Schema } from 'mongoose';
class VoteCastSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }
    /* 
    // @notice An event emitted when a vote has been cast on a proposal
    event VoteCast(
        address voter,
        uint256 proposalId,
        bool support,
        uint256 votes
    );
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
            voter: { type: String },
            support: { type: Boolean, default: false },
            votes: { type: Number, default:0 },
        }, { timestamps: false, strict: false });

        this.objectSchema.index({ contract: 1, chainId: 1, blockNumber: -1, });
    }
}

export default mongoose.model('VoteCast', (new VoteCastSchema()).objectSchema);