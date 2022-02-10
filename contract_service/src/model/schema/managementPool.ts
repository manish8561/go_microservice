import mongoose, { Schema } from 'mongoose';

class ManagementPoolSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            poolType: { type: String },
            status: { type: String, default: 'pending' },
            transactionHash: { type: String },
        }, { timestamps: true, strict: false });
    }
}

export default mongoose.model('ManagementPool', (new ManagementPoolSchema()).objectSchema);