import mongoose, { Schema } from 'mongoose';

class InvestmentRewardLogSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            investmentAmount: { type: String },
            investmentToken: { type: Number },
            time: { type: Number },
            planType: { type: String },
            blockNumber: { type: Number },
            transactionHash: { type: String },
            amtWithdrawn: { type: Boolean, default: false },
        }, { timestamps: false, strict: false });
    }
}

export default mongoose.model('InvestmentRewardLog', (new InvestmentRewardLogSchema()).objectSchema);