import mongoose, { Schema } from 'mongoose';
import eventModel from '../event.model';

class ProposalCanceledSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }
    /* 
    // @notice An event emitted when a proposal has been canceled
    event ProposalCanceled(uint256 id);
    */

    private schema() {
        this.objectSchema = new Schema({
            chainId: { type: Number },
            transactionHash: { type: String },
            blockNumber: { type: Number, default: 0 },
            lastBlockNumber: { type: Number, default: 0 },
            contractName: { type: String, required: true },
            contract: { type: String, required: true },

            proposalId: { type: Number, default: 0 },
        }, { timestamps: false, strict: false });

        this.objectSchema.index({ contract: 1, chainId: 1, blockNumber: -1, });

        this.objectSchema.post('insertMany', async (docs: any) => {
            // calling common function on save for proposal
            eventModel.updateProposalStatus(docs, 'ProposalCanceled');
        });
    }
}

export default mongoose.model('ProposalCanceled', (new ProposalCanceledSchema()).objectSchema);