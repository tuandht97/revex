'use strict';

const config = require('../blockchain/config');
var RealEstate = require('../models/realestate');

const Service = class Service {
    constructor(client) {
        this._client = client;
    }
    async getShareHolder(data) {
        try {
            return await this.query('get_shareholder', data);
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async getRealEstate(id) {
        try {
            return await RealEstate.findOne({ id: id })
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async deleteRealEstate(id) {
        try {
            return await RealEstate.deleteOne({ id });
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listPublishRealEstate() {
        try {
            return await RealEstate.find({ actice: "Publish" });
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async getPublisher(data) {
        try {
            return await this.query('get_publisher', data);
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async login(data) {
        try {
            await this._client.loginUser(data);
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listTransferContractByBuyer(data) {
        try {
            return await this.query('list_transfer_contract_by_buyer', data);
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listTransferContractBySeller(data) {
        try {
            return await this.query('list_transfer_contract_by_seller', data);
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listPublishContract() {
        try {
            return await this.query('list_publish_contract');
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listShareHolder() {
        try {
            return await this.query('list_shareholder');
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listPublisher() {
        try {
            return await this.query('list_publisher');
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listTransferContract() {
        try {
            return await this.query('list_transfer_contract');
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listCode() {
        try {
            return await this.query('list_real_estate');
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    async listSellTritAdvertisingContract() {
        try {
            return await this.query('list_sell_trit_advertising_contract');
        } catch (e) {
            throw (`${e.message}`);
        }
    }
    async listSellTritAdvertisingContractByUser(data) {
        try {
            return await this.query('list_sell_trit_advertising_contract_by_user', data);
        } catch (e) {
            throw (`${e.message}`);
        }
    }

    invoke(fcn, ...args) {
        return this._client.invoke(
            config.chaincodeId, fcn, ...args);
    }

    query(fcn, ...args) {
        return this._client.query(
            config.chaincodeId, fcn, ...args);
    }
};

module.exports.Service = Service;
