/* 
    ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Created          time.Time          `bson:"_created" json:"_created"`
    Modified         time.Time          `bson:"_modified" json:"_modified"`
    Chain_Id         int                `bson:"chain_id" json:"chain_id"`
    Transaction_Hash string             `bson:"transaction_hash" json:"transaction_hash"`
    // Proposal_Type    string             `bson:"proposal_type" json:"proposal_type"`
    Block_Number   int     `bson:"block_number" json:"block_number"`
    Status         string  `bson:"status" json:"status"`
    Proposal_Id    string  `bson:"proposal_id" json:"proposal_id"`
    Proposer       string  `bson:"proposer" json:"proposer"`
    Eta            int     `bson:"eta" json:"eta"`
    Start_Time     int     `bson:"start_time" json:"start_time`
    End_Time       int     `bson:"end_time" json:"end_time`
    Description    string  `bson:"description" json:"description`
    Voting_Period  int     `bson:"voting_period" json:"voting_period` // in days
    For_Votes      float64 `bson:"for_votes" json:"for_votes"`
    Against_Votes  float64 `bson:"against_votes" json:"against_votes"`
    Canceled       bool    `bson:"canceled" json:"canceled"`
    Executed       bool    `bson:"executed" json:"executed"`
    Title          string  `bson:"title" json:"title"`
    Db_Description string  `bson:"db_description" json:"db_description"`
    Proposal_Type  int     `bson:"proposal_type" json:"proposal_type"` //1 for core 2 for community
    Cron_Status    string  `bson:"cron_status" json:"cron_status"`
*/
import mongoose, { Schema } from 'mongoose';
class ProposalSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            _created: { type: Date },
            _modified: { type: Date },
            chain_id: { type: Number },
            transaction_hash: { type: String },
            block_number: { type: Number, default: 0 },
            status: { type: String },

            proposal_id: { type: Number },// proposal id in contract
            proposer: { type: String },
            eta: { type: Number, default: 0 },

            start_time: { type: Number },
            end_time: { type: Number },
            description: { type: String },
            voting_period: { type: Number },
            for_votes: { type: Number },
            against_votes: { type: Number },
            canceled: { type: Boolean, default: false },
            executed: { type: Boolean, default: false },
            title: { type: String },
            db_description: { type: String },
            proposal_type: { type: Number },
            cron_status: { type: String, default: "pending" }

        }, { timestamps: false, strict: false });

        this.objectSchema.index({ contract: 1, chainId: 1, blockNumber: -1, });

        this.objectSchema.post('save', async (doc: any) => {

        });
    }
}

export default mongoose.model('Proposal', (new ProposalSchema()).objectSchema);