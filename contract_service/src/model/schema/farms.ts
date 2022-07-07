import mongoose, { Document, Schema } from 'mongoose';
class FarmSchema extends Schema {
    public objectSchema: any;

    constructor() {
        super();
        this.schema();
    }

    private schema() {
        this.objectSchema = new Schema({
            _created: { type: Date },
            _modified: { type: Date },
            chain_id: { type: Number },
            transaction_hash: { type: String },
            pid: { type: Number },
            address: { type: String },
            name: { type: String },
            token_type: { type: String },
            deposit_token: { type: String },
            status: { type: String },
            masterchef: { type: String },
            router: { type: String },
            weth: { type: String },
            stake: { type: String },
            ac_token: { type: String },
            reward: { type: String },
            rewardImage: { type: String },
            bonus_multiplier: { type: Number },
            token_per_block: { type: Number },
            farmType: { type: String },
            source: { type: String },
            source_link: { type: String },
            autocompound_check: { type: Boolean },
            tvl_staked: { type: Number },// value should be in USD
            daily_apr: { type: Number },
            daily_apy: { type: Number },
            weekly_apy: { type: Number },
            yearly_apy: { type: Number },
            price_pool_token: { type: Number },
            yearly_swap_fees: { type: Number },
            // token0: { type: Object },
            // token1: { type: Object },
            token0: { type: Object },
            token1: { type: Object },
        }, { timestamps: false, strict: false });

        this.objectSchema.index({ address: 1, status: 1, chain_id: 1, name: "text", tvl_staked: 1, token_per_block: 1, token_type: 1 });
    }
}

export default mongoose.model('farms', (new FarmSchema()).objectSchema);