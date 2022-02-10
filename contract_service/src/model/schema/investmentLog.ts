import mongoose, { Schema } from 'mongoose';

class InvestmentLogSchema extends Schema {
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
            ethValue: { type: Number },
            blockNumber: { type: Number },
            transactionHash: { type: String },
            amtWithdrawn: { type: Boolean, default: false },
        }, { timestamps: false, strict: false });

        this.objectSchema.index({ user: 1, investmentAmount: 1, time: 1, blockNumber: 1 });
    }
}

export default mongoose.model('InvestmentLog', (new InvestmentLogSchema()).objectSchema);