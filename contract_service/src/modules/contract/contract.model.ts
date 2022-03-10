import BaseModel from "../../model/base.model";
import masterChefHelper from "./masterChef.helper";
import farms from "../../model/schema/farms";
import { Responses } from "../../helpers";
// import axios from "axios";
// import { initParams } from "request-promise";


class farmModel extends BaseModel {
  constructor() {
    super();
  }
  public async getAprValue(response: any) {
    try {
      const farmData: any = await farms.find({})
      const newArr: any = [];
      if (farmData.length > 0) {
        for (let it of farmData) {
          const { masterchef, deposit_token }: any = it;
          const calApr: any = await masterChefHelper.calculateAPRValue(masterchef, deposit_token);
          it.daily_apr = calApr
          newArr.push(it)
        }
        await farms.insertMany(newArr, { ordered: false }).catch(err => {
          console.error(err);
        })
      }
    }
    catch (error) {
      throw Responses.error(response, { message: error });
    }
  }
}

export default new farmModel();