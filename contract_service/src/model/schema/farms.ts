import mongoose, { Document, Schema } from 'mongoose';
class farmsSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            address: { type: String },
            daily_apr: { type: String },
            daily_apy: { type: String },
            tvl_staked: { type: String },
            chain_Id: { type: Number },
            status: { type: String },
            lastTime: { type: Date },
            cronStatus: { type: String, default: 'no' },
        }, { timestamps: false, strict: false });
        this.objectSchema.index({ cronStatus: 1, })
    }
}

export default mongoose.model('farms', (new farmsSchema()).objectSchema);