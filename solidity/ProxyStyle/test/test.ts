import { ethers } from "hardhat";
import { FacetCutAction, getSelectors } from "./utils";
import { ContractFactory } from "ethers";

describe("Diamond Test", function () {

  const Diamonds = [
    {
      key: 'market',
      data: [
        // Market facets
        'contracts/services/market/facets/Create.sol:Create',
        'contracts/services/market/facets/Get.sol:Get'
      ]
    },

  ];

  let diamondCut: any[] = [];
  let diamondArgs: any = {};

  beforeEach(async function () {
    const accounts = await ethers.getSigners();

    diamondArgs = {
      owner: accounts[0].address,
      init: ethers.ZeroAddress,
      initCalldata: ethers.ZeroAddress
    }

    console.log("Deployer:", diamondArgs.owner, "(owner)");

    diamondCut = await Promise.all(Diamonds.map(async (d, i) => {
      return {
        ...d,
        data: await Promise.all(d.data.map(async (f, j) => {

          const facet = await (await ethers.getContractFactory(f)).deploy();
          console.log(`---------------------------------------------------------------`);
          console.log(`Diamond Name: ${d.key}`);
          console.log(`---------------------------------------------------------------`);
          console.log(`Facet Name: ${f}`);
          console.log(`Deployed Address: ${await facet.getAddress()}`);
          console.log(`Functions: ${await getSelectors(facet)}`);
          console.log(`---------------------------------------------------------------`);

          return {
            facetAddress: await facet.getAddress(),
            action: FacetCutAction.Add,
            functionSelectors: await getSelectors(facet)
          };
        }))
      }
    }))
  })


  it("", async function () {
    const tokens = await ethers.getSigners();

    // console.log(facetCuts, diamondArgs);
    console.log(':: Market Diamond Deploy ::');
    const market: ContractFactory = await ethers.getContractFactory("Market");
    console.log(`---------------------------------------------------------------`);
    console.log(diamondCut);
    console.log(`---------------------------------------------------------------`);
    const marketDiamond: any = await market.deploy(diamondCut, diamondArgs);
    console.log(`---------------------------------------------------------------`);
    console.log('Market: ', await marketDiamond.getAddress());
    console.log(`---------------------------------------------------------------`);

    console.log('Market Diamond Facets:', await marketDiamond.facets());

    describe('', () => {
      it('Market Diamond Loupe Test', async () => {
        const facets = await marketDiamond.facets();
        console.log('Loupe Facet: ', facets[2])
        const functs = facets[2][1];
        for (let i = 0; i < functs.length; i++) {
          console.log(`---------------------------------------------------------------`);
          console.log(await marketDiamond[functs[i]]);
          console.log(`---------------------------------------------------------------`);
        }
      })


      it("New Create", async function () {
        console.log(':: New Create');
        await (await ethers.getContractAt(
          "contracts/services/market/facets/Create.sol:Create",
          await marketDiamond.getAddress())
        ).create(
          await tokens[1].getAddress(),
          await tokens[2].getAddress(),
          1000
        );
      })

    })
  });
});
