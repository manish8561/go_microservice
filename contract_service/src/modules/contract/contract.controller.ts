import { Response, Router } from "express";
import * as Interfaces from "interfaces";
import { Responses } from "../../helpers";
// import { requestDecrypt } from "../../middlewares";
// import Contract from "./contract.model";
// import masterChefAbi from '../../bin/masterChefContractABI.json'
// import masterChefHelper from "./masterChef.helper";
// import farms from "../../model/schema/farms";
import farmModel from './contract.model'

class ContractController implements Interfaces.Controller {
    public path = '/contract';
    public router = Router();
    constructor() {
        this.initializeRoutes();
    }

    private async initializeRoutes() {
        this.router
            .all(`${this.path}/*`)
            .get(
                this.path + "/getapr",
                this.calApy
            );
    }

    private async calApy(req: any, response: Response) {
        try {
            const res: any = await farmModel.getFarmsValue();
            return Responses.success(response, { message: res });;
        } catch (error) {
            return Responses.error(response, { message: error });
        }
    }
}

export default ContractController;