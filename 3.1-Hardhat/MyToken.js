const hre = require("hardhat")
const { expect } = require("chai");
describe("MyToken Test", ( ) => {

    const initialSupply = 10000;
    let MyTokenContract;
    let account1,account2;
    beforeEach (async () => {
        [account1, account2] = await hre.ethers.getSigners();
        const MyToken = await hre.ethers.getContractFactory("ECR721");
        MyTokenContract = await MyToken.deploy(initialSupply); // 部署合约
        MyTokenContract.waitForDeployment();
        const contractAddress = await MyTokenContract.getAddress();
        expect(contractAddress).to.length.greaterThan(0);
        console.log("合约地址:", contractAddress);
        // console.log("等待2秒");
        // await new Promise(resolve => setTimeout(resolve, 2000)); 
        // console.log("开始测试");
    })

    it("验证合约的 name symbol decimals",async () => {
        const name = await MyTokenContract.name();
        const symbol = await MyTokenContract.symbol();
        const decimals = await MyTokenContract.decimals();

        expect(name).to.equal("MyToken");
        expect(symbol).to.equal("MTK");
        expect(decimals).to.equal(18);
        console.log("合约名称:", name);
        console.log("合约符号:", symbol);
        console.log("合约小数位:", decimals);
    })
    it("测试转账",async () => {
        console.log("执行测试2");
        const transferAmount = 1000;
        
        let monney = await MyTokenContract.transfer(account2,transferAmount);
        const balanceBefore = await MyTokenContract.balanceOf(account2);
        console.log("转账金额:", monney)
        expect(transferAmount).to.equal(1000);
    })

})