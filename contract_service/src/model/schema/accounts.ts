import mongoose, { Schema } from "mongoose";

class Accounts extends Schema {
  public mongooseObj: any;

  constructor() {
    super();
    this.schema();
  }

  private schema() {
    const objSchema = new Schema(
      {
        name: { type: String, trim: true },
        role: { type: String, default: 'admin' },
        address: { type: String, trim: true, unique: true },
        referrer: { type: String, trim: true },
        email: { type: String, trim: true, unique: true },
        password: { type: String, trim: true },
        transactionHash: { type: String, trim: true, default: '' },
        status: { type: String, default: 'pending' },
        cronStatus: { type: String, default: 'no' }
      },
      { timestamps: true, strict: false }
    );
    this.mongooseObj = mongoose.model('Account', objSchema);
  }
}

export default (new Accounts()).mongooseObj;
