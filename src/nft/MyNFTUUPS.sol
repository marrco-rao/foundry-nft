// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;
import "@openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol"; 
import "@openzeppelin-contracts-upgradeable/contracts/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-contracts-upgradeable/contracts/token/common/ERC2981Upgradeable.sol";


contract MyNFTUUPS is 
        Initializable,  
        ERC721URIStorageUpgradeable, 
        ERC2981Upgradeable,
        OwnableUpgradeable,
        UUPSUpgradeable
    {

    uint256 public _nextTokenId;
    // 最大供应量
    uint256 public constant MAX_SUPPLY = 10000;

    // 铸造价格
    uint256 public mintPrice = 0.01 ether;

    /**
     * @dev 当新的NFT被成功铸造时触发此事件
     * @param to 接收NFT的地址
     * @param tokenId 新铸造的NFT的唯一标识符
     * @param tokenURI NFT的元数据URI
     */
    event NFTMinted(address indexed to, uint256 indexed tokenId, string tokenURI);

    /**
     * @notice 初始化函数
     * @dev 初始化NFT集合名称和符号，设置合约所有者 
     */
    function initialize(
        string memory name_,
        string memory symbol_,
        address royaltyReceiver,
        uint96 royaltyBps
    ) public initializer {
        __ERC721_init(name_, symbol_);
        __ERC721URIStorage_init();
        __ERC2981_init();
        __Ownable_init(msg.sender);
        
        _setDefaultRoyalty(royaltyReceiver, royaltyBps); // 设置默认版税信息

    }

    // 铸造新的NFT,需要支付铸造费用
    function mint (
        address to,
        string calldata uri
    ) external payable returns (uint256) {
        // 检查是否达到最大供应量
        require(_nextTokenId <= MAX_SUPPLY, "Max supply reached");
        // 检查支付的铸造费用
        require(msg.value >= mintPrice, "Insufficient payment for minting");

        // 生成新的Token ID
        _nextTokenId += 1;
        uint256 newItemId = _nextTokenId; 
        
        // 安全铸造NFT
        _safeMint(to, newItemId);
        // 设置Token URI
        _setTokenURI(newItemId, uri);

        // 触发NFT铸造事件
        emit NFTMinted(to, newItemId, uri);
        return newItemId;
    }

    // 授权UUPS升级
    function _authorizeUpgrade(address) internal override onlyOwner {}

    // 重写supportsInterface
    function supportsInterface(bytes4 interfaceId) 
        public 
        view 
        override(ERC721URIStorageUpgradeable, ERC2981Upgradeable) returns (bool) {
        return super.supportsInterface(interfaceId);
    }


    // 修改NFT的铸造价格
    function setMintPrice(uint256 newPrice) public onlyOwner {
        mintPrice = newPrice;
    }

    // 修改NFT的tokenURI
    function setTokenURI(uint256 tokenId, string memory tokenURI) public onlyOwner {
        _setTokenURI(tokenId, tokenURI);
    }

    // 查询当前铸造的NFT数量
    function totalSupply() public view returns (uint256) {
        return _nextTokenId;    
    }

    // 提取合约中的以太币（铸造费用）到合约所有者地址，仅合约所有者能提取
    /**
     * @dev 提取合约中的以太币（铸造费用）到合约所有者地址
     * @notice 只有合约的所有者可以调用此函数，以确保资金的安全性
    */
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance; 
        require(balance > 0, "No balance to withdraw");   
        (bool success, ) = payable(owner()).call{value: balance}("");
        require(success, "Withdrawal failed");
    }
}
