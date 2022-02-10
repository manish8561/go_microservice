import mongoose, { Schema } from 'mongoose';

class UserBonusSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            hustlerBonus: { type: Number, default: 0 },
            hustlerShare: { type: Number, default: 0 },
            placementBonus: { type: Number, default: 0 },
            pool1: { type: Number, default: 0 },
            pool2: { type: Number, default: 0 },
            pool3: { type: Number, default: 0 },
            oldTeamInvestment: { type: Number, default: 0 },
            worldWideBonus: { type: Number, default: 0 },
            managementBonus: { type: Number, default: 0 },
            lastTime: { type: Date },
            cronStatus: { type: String, default: 'no' },
        }, { timestamps: false, strict: false });
        this.objectSchema.index({ user: 1, cronStatus: 1, })
    }
}

export default mongoose.model('UserBonus', (new UserBonusSchema()).objectSchema);