import { Request, Response, Router } from "express";

import * as Interfaces from "interfaces";
import { Responses } from "../../helpers";
import { requestDecrypt } from "../../middlewares";

import Contract from "./contract.model";


class ContractController implements Interfaces.Controller {
    public path = '/Contract';
    public router = Router();
    constructor() {
        this.initializeRoutes();
    }

    private async initializeRoutes() {
        this.router
            .all(`${this.path}/*`)

            .get(
                this.path + "/getTotalPairs",
                this.getTotalPairs
            )
            .get(
                this.path + "/getTotalLiquidity",
                this.getTotalLiquidity
            )
            .get(
                this.path + "/getTotalVolume",
                this.getTotalVolume
            );
    }

    //Get Contracts
    private async getTotalPairs(req: any, response: Response) {
        if (req.body) {
            
            const chainId = 1280, factory = "0x2A6d5103E4312E487413E476B7aF91FD67be7627";
            
            try {
                const count: string = await Contract.getTotalPairs(chainId, factory);
                return Responses.success(response, { count });
            } catch (error) {
                return Responses.error(response, { message: error });
            }
        }
    }

    private async getTotalLiquidity(req: any, response: Response) {
        if (req.body) {
            
            const chainId = 1280, masterchef = "0x904c75b30B36Ff488FDF44a4980986ddb1595C67",pair = "0x6B7291e1c4c212c411676cAa4499c54A0b634394"  ;
            
            try {
                const liquidity: string = await Contract.getTotalLiquidity(chainId, masterchef, pair);
                return Responses.success(response, { liquidity });
            } catch (error) {
                return Responses.error(response, { message: error });
            }
        }
    }

    private async getTotalVolume(req: any, response: Response) {
        if (req.body) {
            
            const chainId = 1280, masterchef = "0x904c75b30B36Ff488FDF44a4980986ddb1595C67",pair = "0x6B7291e1c4c212c411676cAa4499c54A0b634394"  ;
            
            try {
                const volume: string = await Contract.getTotalLiquidity(chainId, masterchef, pair);
                return Responses.success(response, { volume });
            } catch (error) {
                return Responses.error(response, { message: error });
            }
        }
    }
    
}

export default ContractController;