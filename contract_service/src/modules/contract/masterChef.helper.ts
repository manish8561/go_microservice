import web3Helper from "../../helpers/common/web3.helper";
import pairContractABI from '../../bin/pairContractABI.json'

class MasterChef extends web3Helper {
  constructor() {
    super();
  }




  public async calculateAPRValue(Abi: any, AbiAddress: string): Promise<string> {
    try {
      const contract = await this.callContract(Abi, AbiAddress)
      const tokenPrice = await this.callPrice('0x496b384ee9Cf03Af7d5ED6f1EEbe2c2ba1415242');
      return tokenPrice;
      // return contract;
    } catch (err) {
      throw err;
    }
  }

  public async getTokenZero(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(pairContractABI, pairAddress);
      return await contract.methods.token0().call();
    } catch (error) {
      throw error;
    }
  };

  public async getTokenOne(pairAddress: any): Promise<string> {
    try {
      const contract: any = await this.callContract(pairContractABI, pairAddress);
      return await contract.methods.token1().call();
    } catch (error) {
      throw error;
    }
  };

  public async callPrice(pairAddress: any): Promise<string> {
    try {
      let price = 0;
      if (pairAddress == "0x0000000000000000000000000000000000000000") {
        return '0';
      }

      // console.log("pairAddresspairAddress", pairAddress);

      const tokenZero = await this.getTokenZero(pairAddress);
      const tokenOne = await this.getTokenOne(pairAddress);

      
      console.log("tokenZero", tokenZero)
      console.log("tokenOne", tokenOne)

      return tokenZero;


      // const tokenOne = await ExchangeService.getTokenOne(pairAddress);
      // const reserve = await ExchangeService.getReserves(pairAddress);

      // const decimalZero = await ContractServices.getDecimals(tokenZero);
      // const decimalOne = await ContractServices.getDecimals(tokenOne);

      // console.log(tokenZero, TOKEN_LIST[2].address);

      // if (tokenZero.toLowerCase() === TOKEN_LIST[2].address.toLowerCase()) {
      //     return price = ((reserve[0] * (10 ** decimalOne)) / (reserve[1] * (10 ** decimalZero)));
      // }

      // if (tokenOne.toLowerCase() === TOKEN_LIST[2].address.toLowerCase()) {
      //     return price = ((reserve[1] * (10 ** decimalZero)) / (reserve[0] * (10 ** decimalOne)));
      // }

      // let priceBNBToUSD = calPrice(BNB_BUSD_LP); //replace with BNB-USD pair

      // if (tokenZero.toLowerCase() === WETH.toLowerCase()) {
      //     price = ((reserve[0] * (10 ** decimalOne)) / (reserve[1] * (10 ** decimalZero)));
      //     return (price * priceBNBToUSD);
      // }

      // if (tokenOne.toLowerCase() === WETH.toLowerCase()) {
      //     price = ((reserve[1] * (10 ** decimalZero)) / (reserve[0] * (10 ** decimalOne)));
      //     return (price * priceBNBToUSD);
      // }
      // return 0
    } catch (err) {
      throw err
    }
  }

}

export default new MasterChef();
