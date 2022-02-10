import mongoose, { Schema } from 'mongoose';

class ManagementPoolRegistrationSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            user: { type: String },
            poolType: { type: String },
            time: { type: Number },
            blockNumber: { type: Number },
            transactionHash: { type: String },
        }, { timestamps: false, strict: false });
    }
}

export default mongoose.model('ManagementPoolRegistration', (new ManagementPoolRegistrationSchema()).objectSchema);