// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// IERC1820Registry 接口定义
interface IERC1820Registry {
    function setInterfaceImplementer(
        address account,
        bytes32 interfaceHash,
        address implementer
    ) external;

    function getInterfaceImplementer(
        address account,
        bytes32 interfaceHash
    ) external view returns (address);
}

// IERC777 接口定义
interface IERC777 {
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function granularity() external view returns (uint256);
    function balanceOf(address tokenHolder) external view returns (uint256);
    function totalSupply() external view returns (uint256);
    function send(address recipient, uint256 amount, bytes calldata data) external;
    function burn(uint256 amount, bytes calldata data) external;
    function isOperatorFor(address operator, address tokenHolder) external view returns (bool);
    function authorizeOperator(address operator) external;
    function revokeOperator(address operator) external;
    function defaultOperators() external pure returns (address[] memory);
    function operatorSend(
        address sender,
        address recipient,
        uint256 amount,
        bytes calldata data,
        address operator
    ) external;

    event Sent(
        address indexed operator,
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data,
        bytes operatorData
    );
    event Minted(
        address indexed operator,
        address indexed to,
        uint256 amount,
        bytes data,
        bytes operatorData
    );
    event Burned(
        address indexed operator,
        address indexed from,
        uint256 amount,
        bytes data,
        bytes operatorData
    );
    event AuthorizedOperator(
        address indexed operator,
        address indexed tokenHolder
    );
    event RevokedOperator(
        address indexed operator,
        address indexed tokenHolder
    );
}

// IERC777Recipient 接口定义
interface IERC777Recipient {
    function tokensReceived(
        address operator,
        address from,
        uint256 amount,
        bytes calldata data,
        bytes calldata operatorData
    ) external;
}

// ERC777 代币合约
contract MyToken is IERC777, IERC777Recipient {

    string private _name;
    string private _symbol;
    uint256 private _granularity;
    uint256 private _totalSupply;
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => bool)) private _operators;
    address private constant ERC1820_REGISTRY = 0x1820a4B7618BdE71Dce8cdc73aAB6C95905faD24;
    bytes32 private constant ERC777_INTERFACE_ID = keccak256("ERC777Token");

    constructor(string memory name_, string memory symbol_, uint256 initialSupply_) {
        _name = name_;
        _symbol = symbol_;
        _granularity = 1;

        // 注册 ERC1820 接口
        IERC1820Registry(ERC1820_REGISTRY).setInterfaceImplementer(
            address(this),
            ERC777_INTERFACE_ID,
            address(this)
        );

        // 初始发行
        _mint(msg.sender, initialSupply_);
    }

    // ERC777 标准函数实现
    function name() public view override returns (string memory) {
        return _name;
    }

    function symbol() public view override returns (string memory) {
        return _symbol;
    }

    function granularity() public view override returns (uint256) {
        return _granularity;
    }

    function balanceOf(address tokenHolder) public view override returns (uint256) {
        return _balances[tokenHolder];
    }

    function totalSupply() public view override returns (uint256) {
        return _totalSupply;
    }

    function send(address recipient, uint256 amount, bytes calldata data) external override {
        _send(msg.sender, recipient, amount, data, msg.sender);
    }

    function operatorSend(
        address sender,
        address recipient,
        uint256 amount,
        bytes calldata data,
        address operator
    ) external override {
        require(isOperatorFor(operator, sender), "Not an operator");
        _send(sender, recipient, amount, data, operator);
    }

    function isOperatorFor(address operator, address tokenHolder) public view override returns (bool) {
        return _operators[tokenHolder][operator];
    }

    function authorizeOperator(address operator) external override {
        _operators[msg.sender][operator] = true;
        emit AuthorizedOperator(operator, msg.sender);
    }

    function revokeOperator(address operator) external override {
        _operators[msg.sender][operator] = false;
        emit RevokedOperator(operator, msg.sender);
    }

    function defaultOperators() external pure override returns (address[] memory) {
        return new address[](0);
    }

    function burn(uint256 amount, bytes calldata data) external override {
        _burn(msg.sender, amount, data);
    }

    // 内部函数
    function _send(
        address from,
        address to,
        uint256 amount,
        bytes memory data,
        address operator
    ) internal {
        require(from != address(0), "Send from zero address");
        require(to != address(0), "Send to zero address");
        require(_balances[from] >= amount, "Insufficient balance");

        // 扣除发送方余额
        _balances[from] -= amount;
        // 增加接收方余额
        _balances[to] += amount;

        // 触发 Sent 事件
        emit Sent(operator, from, to, amount, data, "");

        // 如果接收方是合约，调用 tokensReceived 钩子
        if (to.code.length > 0) {
            try IERC777Recipient(to).tokensReceived(operator, from, amount, data, "") {} catch {}
        }
    }

    function _burn(
        address from,
        uint256 amount,
        bytes memory data
    ) internal {
        require(_balances[from] >= amount, "Insufficient balance");

        _balances[from] -= amount;
        _totalSupply -= amount;

        emit Burned(msg.sender, from, amount, data, "");
    }

    function _mint(
        address to,
        uint256 amount
    ) internal {
        require(to != address(0), "Mint to zero address");

        _balances[to] += amount;
        _totalSupply += amount;

        emit Minted(msg.sender, to, amount, "", "");
    }

    // tokensReceived 钩子函数
    function tokensReceived(
        address operator,
        address from,
        uint256 amount,
        bytes calldata data,
        bytes calldata operatorData
    ) external override {
        require(msg.sender == address(this), "Only this token can call this hook");
        // 自动处理逻辑：例如记录日志或触发其他操作
        emit TokensReceivedHook(from, amount);
    }

    event TokensReceivedHook(address indexed from, uint256 amount);
}

// 操作员合约
contract MyOperator {
    MyToken public immutable token;

    constructor(address tokenAddress) {
        token = MyToken(tokenAddress);
    }

    function sendOnBehalf(
        address sender,
        address recipient,
        uint256 amount,
        bytes calldata data
    ) external {
        require(token.isOperatorFor(address(this), sender), "Not authorized operator");
        token.operatorSend(sender, recipient, amount, data, address(this));
    }
}