import mongoose, { Schema } from 'mongoose';

class LevelIncomeSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            referrer: { type: String },
            level: { type: Number },
            time: { type: Number },
            value: { type: Number },
            investment: { type: Number },
            blockNumber: { type: Number },
            transactionHash: { type: String },
            cronStatus: { type: String, default: 'no' },
        }, { timestamps: false, strict: false });

        this.objectSchema.index({ user: 1, referrer: 1, value: 1, investment: 1, time: 1, blockNumber: 1 });
    }
}

export default mongoose.model('LevelIncome', (new LevelIncomeSchema()).objectSchema);