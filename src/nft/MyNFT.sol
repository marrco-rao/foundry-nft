// SPDX-License-Identifier: MIT
pragma solidity ^0.8.29;
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract MyNFT is ERC721URIStorage, Ownable {
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
     * 
     * 此事件用于记录NFT的铸造操作，包括接收者地址、NFT资产ID和对应的元数据URI。
     * 外部应用和用户界面可以监听此事件以实时获取NFT铸造的信息。
     */
    event NFTMinted(address indexed to, uint256 indexed tokenId, string tokenURI);

    /**
     * @notice 构造函数
     * @dev 初始化NFT集合名称和符号，设置合约所有者 
     */
    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender) {}

    // 铸造新的NFT,需要支付铸造费用
    function mintNFT(
        address to,
        string calldata tokenURI
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
        _setTokenURI(newItemId, tokenURI);

        // 触发NFT铸造事件
        emit NFTMinted(to, newItemId, tokenURI);
        return newItemId;
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
        payable(owner()).transfer(balance);
    }

    // NFT资产销毁
    function burn(uint256 tokenId) external  {
        // 校验NFT的所有权或合约所有者权限
        require(ownerOf(tokenId) == msg.sender || msg.sender == owner(), "Not authorized");
        _burn(tokenId);
    }
    // 重写supportsInterface
    function supportsInterface(bytes4 interfaceId) public view override(ERC721URIStorage) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
