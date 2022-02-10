/* common file to add all controllers(modules) in the project */
import ContractController from "./contract/contract.controller";
import PriceFeedController from "./pricefeed/pricefeed.controller";

export default [new PriceFeedController(), new ContractController()];