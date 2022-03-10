import mongoose, { Document, Schema } from 'mongoose';
class farmsSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super()
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            _created: { type: Date },
            _modified: { type: Date },
            chain_Id: { type: Number },
            transaction_hash: { type: String },
            pid: { type: Number },
            address: { type: String },
            name: { type: String },
            token_type:{type: String},
            deposit_token:{type: String},
            masterchef:{type: String},



            tvl_staked: { type: Number },// value should be in USD
            daily_apr: { type: Number },
            daily_apy: { type: Number },
            status: { type: String },




            token0: {type:Object},
            token1:{type:Object},

	Router             string`bson:"router" json:"router"`
	Weth               string`bson:"weth" json:"weth"`
	Stake              string`bson:"stake" json:"stake"`       //staking contract address
	AC_Token           string`bson:"ac_token" json:"ac_token"` //autocompound token
	Reward             string`bson:"reward" json:"reward"`     //cake address
	Bonus_Multiplier   int`bson:"bonus_multiplier" json:"bonus_multiplier"`
	Token_Per_Block    int`bson:"token_per_block" json:"token_per_block"`
	Source             string`bson:"source" json:"source"`
	Source_Link        string`bson:"source_link" json:"source_link"`
	Autocompound_Check bool`bson:"autocompound_check" json:"autocompound_check"`
	Weekly_APY         float64`bson:"weekly_apy" json:"weekly_apy"`
	Yearly_APY         float64`bson:"yearly_apy" json:"yearly_apy"`
	Price_Pool_Token   float64`bson:"price_pool_token" json:"price_pool_token"`
	Yearly_Swap_Fees   float64`bson:"yearly_swap_fees" json:"yearly_swap_fees"`
	Gauge_Info         string`bson:"gauge_info" json:"gauge_info"`
        }, { timestamps: false, strict: false });
this.objectSchema.index({  address: 1, status:1 })
    }
}

export default mongoose.model('farms', (new farmsSchema()).objectSchema);