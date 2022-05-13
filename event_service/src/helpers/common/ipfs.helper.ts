import * as requestPromise from 'request-promise';

class IPFSHelper {
    // private ipfs_url = "https://ipfs.infura.io:5001/api/v0/";
    constructor() { }
    /**
     * read file from ipfs
     * since it is json object in string form
     * @param  {string} hash
     * @returns Promise
     */
    public async readFile(hash: string): Promise<any> {
        try {
            const ipfs_url = process.env.INFURA_IPFS_URL;
            const r = await requestPromise.get(`${ipfs_url}get?arg=${hash}`);

            return JSON.parse(r);
        } catch (error) {
            throw error;
        }
    }
}
export default new IPFSHelper();