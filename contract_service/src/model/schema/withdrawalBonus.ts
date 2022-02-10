import mongoose, { Schema } from 'mongoose';

class WithdrawalBonusSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            time: { type: Number },
            value: { type: Number },
            usdValue: { type: String },
            blockNumber: { type: Number },
            transactionHash: { type: String },

        }, { timestamps: false, strict: false });
        this.objectSchema.index({ time: -1, user: 1 });
    }
}

export default mongoose.model('WithdrawalBonus', (new WithdrawalBonusSchema()).objectSchema);