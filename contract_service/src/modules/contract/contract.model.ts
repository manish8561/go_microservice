import BaseModel from "../../model/base.model";
import masterChefHelper from "./masterChef.helper";
import farms from "../../model/schema/farms";
import { Responses } from "../../helpers";
import axios from "axios";


class farmModel extends BaseModel {
  constructor() {
    super();
  }


  public async getAprValue(response: any) {
    const farmData: any = await axios.get(`${process.env.FARM_API_URL}farm?page=1&limit=10&status=active`);

    if (farmData.status === 200 && farmData.data.data.length > 0) {
      try {

        if (farmData && farmData.data.data.length > 0) {
          let newData: any = [];
          for (let it  of farmData.data.data){
            const { masterchef, deposit_token }: any = it;
            // console.log("masterChecfffffffff", masterchef, deposit_token)
            const calApr: any = await masterChefHelper.calculateAPRValue(masterchef, deposit_token);
            it.daily_apr = calApr
            // if (it.daily_apr) {
            //   console.log("april", calApr)

            //   it.daily_apr = calApr
            // }
            // console.log("ittttttttttttt", it)

            newData.push(it);
            console.log("newDataaaaaaaaaaa1", newData)

          }
        
          console.log("newDataaaaaaaaaaa2", newData)

        }


        // await farms.insertMany(farmData.data.data, { ordered: false }).catch(err => {
        //   console.error(err);
        // })



        const farmList = await farms.find({});

        // if (farmList && farmList.length > 0) {
        //   farmList.map((it: any) => {
        //     const mC : any = it.masterchef
        //     console.log('HERE', it, i)
        //   })
        // }

        // console.log("FARMMMMMMMMMMMMMMM", farmList);
        // const calApr: any = await masterChefHelper.calculateAPRValue(abi, ContarctAddess);
        // const data: any = await new farms({ apr: calApr }).save();
        return farmList
      } catch (error) {
        throw Responses.error(response, { message: error });
      }

    }


    // }
  }


}

export default new farmModel();