// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {DiamondFacade} from "../../../modules/diamond/DiamondFacade.sol";
import {IOrderbook} from "./IOrderbook.sol";
import {Data} from "./Data.sol";

contract Orderbook is DiamondFacade {
    using Data for Data.Storage;

    Data.Storage internal $;

    constructor(
        address _base,
        address _quote,
        uint _price,
        address _app
    ) DiamondFacade("orderbook", _app) {

        $.base = _base;
        $.quote = _quote;
        $.price = _price;
    }
}
