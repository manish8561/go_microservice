import mongoose, { Schema } from 'mongoose';

class RegistrationSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            referrer: { type: String },
            email: { type: String },
            referrerEmail: { type: String },
            time: { type: Number },
            blockNumber: { type: Number },
            transactionHash: { type: String },
            teamInvestment: { type: Number, default: 0 },
            currentRank: { type: String, default: 'none' },
            commissionBonus: { type: Number, default: 0 },
            closestCommission: { type: Number, default: 0 },
            shareWorldpool: { type: Number, default: 0 },
            teamInvestmentManagement: { type: Number, default: 0 },

        }, { timestamps: false, strict: false });
        this.objectSchema.index({ user: 1, referrer: 1, email: 1, time: -1, blockNumber: 1, teamInvestment:1 })
    }
}

export default mongoose.model('Registration', (new RegistrationSchema()).objectSchema);