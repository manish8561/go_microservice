import BaseModel from "../../model/base.model";
import request from "request";
class PriceFeed extends BaseModel {
    constructor() {
        super();
    }

    public async getPrice(id: string): Promise<any> {
        // console.log(symbol, usd);
        var options = {
            url: `https://api.coingecko.com/api/v3/simple/price?ids=${id}&vs_currencies=usd`,
        };
        return new Promise((resolve, reject) => {
            request(options, async (error, res, body) => {
                if (!error && res.statusCode == 200) {
                    resolve(JSON.parse(body));
                } else {
                    reject(error);
                }
            });
        });
    }


}

export default new PriceFeed();