import mongoose, { Schema } from 'mongoose';

class MaturityWithdrawalSchema extends Schema {
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
            index:{type:String},
            blockNumber: { type: Number},
            transactionHash: { type: String },
       
        }, { timestamps: false, strict: false });
        this.objectSchema.index({time:-1, user:1});
    }
}

export default mongoose.model('MaturityWithdrawal', (new MaturityWithdrawalSchema()).objectSchema);