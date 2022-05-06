import mongoose, { Schema } from 'mongoose';
import eventModel from '../event.model';
class ProposalQueuedSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }
    /* 
        // @notice An event emitted when a proposal has been queued in the Timelock
    event ProposalQueued(uint256 id, uint256 eta);
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
            eta: { type: Number, default: 0 },
        }, { timestamps: false, strict: false });

        this.objectSchema.index({ contract: 1, chainId: 1, blockNumber: -1, });

        this.objectSchema.post('insertMany', async (docs: any) => {
            // calling common function on save for proposal
            eventModel.updateProposalStatus(docs, 'ProposalQueued');
        });
    }
}

export default mongoose.model('ProposalQueued', (new ProposalQueuedSchema()).objectSchema);