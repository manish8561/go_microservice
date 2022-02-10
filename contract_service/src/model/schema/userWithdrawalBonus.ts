import mongoose, { Schema } from 'mongoose';

class UserWithdrawalBonusSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            value: { type: Number },
            usdValue: { type: Number },
            blockNumber: { type: Number },
            transactionHash: { type: String },
            status: { type: String, default: 'pending' },
            message: { type: String, default: '' },

        }, { timestamps: true, strict: false });
        this.objectSchema.index({ transactionHash: -1, user: 1, usdValue: 1 });
    }
}

export default mongoose.model('UserWithdrawalBonus', (new UserWithdrawalBonusSchema()).objectSchema);