// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "../../contracts/access/Ownable.sol";

import "../../contracts/utils/Counters.sol";

import "../../contracts/utils/math/SafeMath.sol";

import "../../contracts/token/ERC20/IERC20.sol";


abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}


library Model {

    struct GameInfo {

        // ID

        uint256 id;

        // A队

        string a_team;

        // B队

        string b_team;

        // A队logo

        string a_team_logo;

        // B队logo

        string b_team_logo;

        // 标签

        string label;

        // 投注开始时间

        uint64 bet_start_time;

        // 投注截止时间

        uint64 bet_end_time;

        // 比赛开始时间

        uint64 game_start_time;

        // 总共投注代币数量

        uint256 total_bet_amount;

        // 胜投注代币数量

        uint256 win_bet_amount;

        // 负投注代币数量

        uint256 fail_bet_amount;

        // 平投注代币数量

        uint256 draw_bet_amount;

        // 胜投注的赔率

        uint256 win_odds;

        // 负投注的赔率

        uint256 fail_odds;

        // 平投注的赔率

        uint256 draw_odds;

        // 比赛是否结束

        bool is_end;

        //比赛结果

        uint8 rst;

        // A队得分

        uint256 a_score;

        // B队得分

        uint256 b_score;

        // 预计平台分红

        uint256 system_dividend;

        // 投注信息

        BetInfo[] bet_infos;

        // 领取日志

        WithDrawLog[] with_draw_logs;

        // 平台是否收获

        bool is_hav;

    }



    struct BetInfo {

        // 投注地址

        address addr;

        // 投注数量

        uint256 amount;

        // 投注类型

        uint8 t;

    }



    struct WithDrawLog {

        // 领取地址

        address addr;

        // 领取数量

        uint256 amount;

    }

}



contract Cup is Ownable {

    // 质押token地址

    IERC20 STAKE_TOKEN;

    // 平台收益地址

    address PLATFORM_ADDRESS;

    // 平台分红比例

    uint256 public SYSTEM_DIVIDEND_RATIO = 10;

    // 最小下注金额

    uint256 public MIN_BET_AMOUNT = 1 ether;

    // 枚举 胜0,负1,平2

    uint8 constant WIN = 0;

    uint8 constant FAIL = 1;

    uint8 constant DRAW = 2;



    // 自增游戏ID

    using Counters for Counters.Counter;

    Counters.Counter private _gameIds;

    mapping(uint256 => Model.GameInfo) private _game;



    constructor(address _token_addr, address _platform_addr) {

        STAKE_TOKEN = IERC20(_token_addr);

        STAKE_TOKEN.approve(address(this), 1000000000 ether);

        PLATFORM_ADDRESS = _platform_addr;

    }



    // 创建游戏

    function createGame(

        string memory a_team,

        string memory b_team,

        string memory a_team_logo,

        string memory b_team_logo,

        string memory label,

        uint64 bet_start_time,

        uint64 bet_end_time,

        uint64 game_start_time

    ) public returns (uint256) {

        require(owner() == _msgSender(), "Ownable: caller is not the owner");

        require(

            bet_start_time < bet_end_time && bet_end_time < game_start_time,

            "time invalid"

        );



        _gameIds.increment();

        uint256 gameId = _gameIds.current();

        _game[gameId].id = gameId;

        _game[gameId].a_team = a_team;

        _game[gameId].b_team = b_team;

        _game[gameId].a_team_logo = a_team_logo;

        _game[gameId].b_team_logo = b_team_logo;

        _game[gameId].label = label;

        _game[gameId].bet_start_time = bet_start_time;

        _game[gameId].bet_end_time = bet_end_time;

        _game[gameId].game_start_time = game_start_time;

        return gameId;

    }



    // 修改游戏

    function updateGame(

        uint256 gameId,

        string memory a_team,

        string memory b_team,

        string memory a_team_logo,

        string memory b_team_logo,

        string memory label,

        uint64 bet_start_time,

        uint64 bet_end_time,

        uint64 game_start_time

    ) external {

        require(owner() == _msgSender(), "Ownable: caller is not the owner");

        require(_game[gameId].id != 0, "no data");

        require(

            bet_start_time < bet_end_time && bet_end_time < game_start_time,

            "time invalid"

        );

        _game[gameId].a_team = a_team;

        _game[gameId].b_team = b_team;

        _game[gameId].a_team_logo = a_team_logo;

        _game[gameId].b_team_logo = b_team_logo;

        _game[gameId].label = label;

        _game[gameId].bet_start_time = bet_start_time;

        _game[gameId].bet_end_time = bet_end_time;

        _game[gameId].game_start_time = game_start_time;

    }



    // 公布结果

    function announce(

        uint256 gameId,

        uint256 a_score,

        uint256 b_score

    ) public {

        require(owner() == _msgSender(), "Ownable: caller is not the owner");

        require(_game[gameId].id != 0, "no data");

        require(_game[gameId].is_end == false, "this game is end");

        _game[gameId].a_score = a_score;

        _game[gameId].b_score = b_score;

        if (a_score == b_score) {

            _game[gameId].rst = DRAW;

        }

        if (a_score > b_score) {

            _game[gameId].rst = WIN;

        }

        if (a_score < b_score) {

            _game[gameId].rst = FAIL;

        }

        _game[gameId].is_end = true;

    }



    // 修改平台分红比例

    function setSystemDividendRatio(uint256 value) public returns (bool) {

        require(owner() == _msgSender(), "Ownable: caller is not the owner");

        require(value > 0 && value < 100, "value should be between 0 and 100");

        SYSTEM_DIVIDEND_RATIO = value;

        return true;

    }



    // 修改最小下注金额

    function setMinBetAmount(uint256 value) public returns (bool) {

        require(owner() == _msgSender(), "Ownable: caller is not the owner");

        require(value > 0, "value should be more than 0");

        MIN_BET_AMOUNT = value;

        return true;

    }



    // 投注

    function bet(

        uint256 gameId,

        uint256 amount,

        uint8 t

    ) external {

        require(_game[gameId].id != 0, "no data");

        require(amount >= MIN_BET_AMOUNT, "amount more than min bet amount");

        require(_game[gameId].is_end == false, "this game bet is end");

        STAKE_TOKEN.transferFrom(_msgSender(), address(this), amount);



        if (t == WIN) {

            uint256 _win_old = _game[gameId].win_bet_amount;

            (, _game[gameId].win_bet_amount) = SafeMath.tryAdd(

                _win_old,

                amount

            );

        }

        if (t == FAIL) {

            uint256 _fail_old = _game[gameId].fail_bet_amount;

            (, _game[gameId].fail_bet_amount) = SafeMath.tryAdd(

                _fail_old,

                amount

            );

        }

        if (t == DRAW) {

            uint256 _draw_old = _game[gameId].draw_bet_amount;

            (, _game[gameId].draw_bet_amount) = SafeMath.tryAdd(

                _draw_old,

                amount

            );

        }

        uint256 _total_old = _game[gameId].total_bet_amount;

        (, _game[gameId].total_bet_amount) = SafeMath.tryAdd(

            _total_old,

            amount

        );



        Model.BetInfo memory _bet_info = Model.BetInfo({

        addr: _msgSender(),

        amount: amount,

        t: t

        });

        _game[gameId].bet_infos.push(_bet_info);



        //计算平台分红

        uint256 _mul_value;

        (, _mul_value) = SafeMath.tryMul(

            _game[gameId].total_bet_amount,

            SYSTEM_DIVIDEND_RATIO

        );

        (, _game[gameId].system_dividend) = SafeMath.tryDiv(_mul_value, 100);



        // 计算赔率

        _odds(gameId);

    }



    // 获取游戏

    function getGame(uint256 gameId)

    public

    view

    virtual

    returns (Model.GameInfo memory)

    {

        require(_game[gameId].id != 0, "no data");

        return _game[gameId];

    }



    // 奖金瓜分的金额

    function getDivideAmount(uint256 gameId)

    public

    view

    virtual

    returns (uint256)

    {

        return _getDivideAmount(gameId);

    }



    // 奖金瓜分的金额

    function _getDivideAmount(uint256 gameId) internal view returns (uint256) {

        require(_game[gameId].id != 0, "no data");

        uint256 _amount;

        (, _amount) = SafeMath.tryMul(

            _game[gameId].total_bet_amount,

            100 - SYSTEM_DIVIDEND_RATIO

        );

        return _amount;

    }



    // 奖金

    function harvest(uint256 gameId, uint8 t)

    public

    view

    virtual

    returns (uint256)

    {

        return _harvest(gameId, _msgSender(), t);

    }



    // 奖金

    function _harvest(

        uint256 gameId,

        address _address,

        uint8 t

    ) internal view returns (uint256) {

        require(_game[gameId].id != 0, "no data");

        uint256 _d_value = _shareRatio(gameId, _address, t);

        return

        SafeMath.mul(

            SafeMath.div(_getDivideAmount(gameId), 1000000),

            _d_value

        );

    }



    // 单用户在某方向投入总金额

    function totalAmountByType(uint256 gameId, uint8 t)

    public

    view

    virtual

    returns (uint256)

    {

        return _totalAmountByType(gameId, _msgSender(), t);

    }



    // 单用户在某方向投入总金额

    function _totalAmountByType(

        uint256 gameId,

        address _address,

        uint8 t

    ) internal view returns (uint256) {

        require(_game[gameId].id != 0, "no data");

        uint256 _b_amount;

        for (uint256 i = 0; i < _game[gameId].bet_infos.length; i++) {

            if (

                _game[gameId].bet_infos[i].t == t &&

                _game[gameId].bet_infos[i].addr == _address

            ) {

                uint256 _temp = _game[gameId].bet_infos[i].amount;

                (, _b_amount) = SafeMath.tryAdd(_b_amount, _temp);

            }

        }

        return _b_amount;

    }



    // 单用户的瓜分比例

    function shareRatio(uint256 gameId, uint8 t)

    public

    view

    virtual

    returns (uint256)

    {

        return _shareRatio(gameId, _msgSender(), t);

    }



    // 单用户的瓜分比例

    function _shareRatio(

        uint256 gameId,

        address _address,

        uint8 t

    ) internal view returns (uint256) {

        require(_game[gameId].id != 0, "no data");

        uint256 _b_amount = _totalAmountByType(gameId, _address, t);

        uint256 _c_amount;

        if (t == WIN) {

            _c_amount = _game[gameId].win_bet_amount;

        }

        if (t == FAIL) {

            _c_amount = _game[gameId].fail_bet_amount;

        }

        if (t == DRAW) {

            _c_amount = _game[gameId].draw_bet_amount;

        }

        return SafeMath.div(SafeMath.mul(_b_amount, 1000000), _c_amount);

    }



    // 计算赔率

    function _odds(uint256 gameId) internal {

        uint256 _amount = _getDivideAmount(gameId);

        if (_game[gameId].win_bet_amount > 0) {

            _game[gameId].win_odds = SafeMath.div(

                SafeMath.mul(_amount, 100),

                _game[gameId].win_bet_amount

            );

        }



        if (_game[gameId].fail_bet_amount > 0) {

            _game[gameId].fail_odds = SafeMath.div(

                SafeMath.mul(_amount, 100),

                _game[gameId].fail_bet_amount

            );

        }



        if (_game[gameId].draw_bet_amount > 0) {

            _game[gameId].draw_odds = SafeMath.div(

                SafeMath.mul(_amount, 100),

                _game[gameId].draw_bet_amount

            );

        }

    }



    // 平台提现

    function platformWithdraw(uint256 gameId) external {

        require(_game[gameId].id != 0, "no data");

        require(_game[gameId].is_end == true, "this game bet is end");

        require(_game[gameId].is_hav == false, "this game is harvest");

        require(

            _msgSender() == PLATFORM_ADDRESS,

            "caller is not the PLATFORM_ADDRESS"

        );

        if (_game[gameId].bet_infos.length < 2) {

            return;

        }

        // 平台收益

        uint256 amount = _game[gameId].system_dividend;

        if (amount > 0) {

            STAKE_TOKEN.transferFrom(address(this), _msgSender(), amount);

            _game[gameId].is_hav = true;

        }

    }



    // 提现

    function withdraw(uint256 gameId) external {

        require(_game[gameId].id != 0, "no data");

        for (uint256 i = 0; i < _game[gameId].with_draw_logs.length; i++) {

            if (_game[gameId].with_draw_logs[i].addr == _msgSender()) {

                return;

            }

        }

        // 只有一个投注

        if (_game[gameId].bet_infos.length == 1) {

            if (_game[gameId].bet_infos[0].t == _game[gameId].rst) {

                STAKE_TOKEN.transferFrom(

                    address(this),

                    _msgSender(),

                    _game[gameId].bet_infos[0].amount

                );

                Model.WithDrawLog memory log = Model.WithDrawLog({

                addr: _msgSender(),

                amount: _game[gameId].bet_infos[0].amount

                });

                _game[gameId].with_draw_logs.push(log);

                return;

            }

        }

        uint256 amount = _harvest(gameId, _msgSender(), _game[gameId].rst);

        if (amount > 0) {

            STAKE_TOKEN.transferFrom(address(this), _msgSender(), amount);

            Model.WithDrawLog memory log = Model.WithDrawLog({

            addr: _msgSender(),

            amount: amount

            });

            _game[gameId].with_draw_logs.push(log);

        }

    }

}

