// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../lib/data_types.sol";
import "../lib/const.sol";
import "../lib/game_operate.sol";
// ownership
import {OwnShip} from "../base_contract/ownship.sol";

// erc20 interface;
import {IERC20} from "../openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

contract WorldCup is OwnShip {
    mapping(uint256 => DataDefine.OneGameInfo) batch_game;
    uint256 public now_game_id = 0;

    modifier LastGameOver() {
        require(
            batch_game[now_game_id].game_info.game_status ==
                DataDefine.GameStatus.kEnd,
            "last game not over"
        );
        _;
    }

    modifier CheckStartTimeAndEndTime() {
        require(
            batch_game[now_game_id].game_info.start_time < block.timestamp &&
                batch_game[now_game_id].game_info.end_time > block.timestamp,
            "time invalid"
        );
        _;
    }

    modifier CheckCanGetReward() {
        DataDefine.UserInfo memory tmp_user_info = GameOperate.GetGameUserInfo(
            batch_game[now_game_id],
            msg.sender
        );

        require(
            tmp_user_info.choose_result ==
                GameOperate.GetResultByGameInfo(batch_game[now_game_id]),
            "not win"
        );

        require(!tmp_user_info.got_reward, "got reward already");
        _;
    }

    modifier CheckGameNotEnd() {
        require(
            GameOperate.GetGameStatus(batch_game[now_game_id].game_info) !=
                DataDefine.GameStatus.kEnd,
            "game is over"
        );

        _;
    }

    modifier CheckGameNotStart() {
        require(
            GameOperate.GetGameStatus(batch_game[now_game_id].game_info) ==
                DataDefine.GameStatus.kNoStart,
            "game is started"
        );

        _;
    }

    modifier CheckGameStart() {
        require(
            GameOperate.GetGameStatus(batch_game[now_game_id].game_info) ==
                DataDefine.GameStatus.kIng,
            "game is not start"
        );

        _;
    }

    function NewGame(
        string[] memory team_infos,
        uint256[] memory time_infos,
        uint256[] memory fee_infos,
        address stake_address
    ) public LastGameOver {
        require(team_infos.length == Const.kStructLen, "team_infos invalid");
        require(
            (time_infos.length == Const.kStructLen &&
                time_infos[Const.kFirst] > time_infos[Const.kSecond]),
            "time_infos invalid"
        );

        require(
            (fee_infos.length == Const.kStructLen &&
                fee_infos[Const.kSecond] > fee_infos[Const.kFirst]),
            "fee_infos invalid"
        );

        // for add one
        uint256 new_game_id = ++now_game_id;

        GameOperate.InitGameInfo(
            batch_game[new_game_id].game_info,
            team_infos,
            time_infos,
            stake_address
        );
        GameOperate.InitFeeInfos(batch_game[new_game_id].fee_info, fee_infos);
    }

    function StartNowGame() public OnlyOwner CheckGameNotStart {
        // start now_game_id game
        GameOperate.StartGame(batch_game[now_game_id].game_info);
    }

    // we can manu stop the game before
    // set game result
    function StopNowGame(uint16[2] memory scores)
        public
        OnlyOwner
        CheckGameNotEnd
    {
        GameOperate.SetGameScores(scores, batch_game[now_game_id].game_info);
        GameOperate.StopGame(batch_game[now_game_id].game_info);
    }

    function GetUserInfosByAddress(address user_address, uint256 game_id)
        public
        view
        returns (DataDefine.UserInfo memory)
    {
        return GameOperate.GetGameUserInfo(batch_game[game_id], user_address);
    }

    function StakeForGame(uint256 amount, DataDefine.GameResult user_choose)
        public
        CheckGameStart
        CheckStartTimeAndEndTime
    {
        // TODO(kingxinwang): add transcation
        require(
            IERC20(GameOperate.GetRequireTokenAddress(batch_game[now_game_id]))
                .transferFrom(msg.sender, address(this), amount),
            "transfer err"
        );

        DataDefine.UserInfo memory tmp_user_info = DataDefine.UserInfo({
            amount: amount,
            choose_result: user_choose,
            got_reward: false
        });

        GameOperate.AddGameStakeUser(
            batch_game[now_game_id],
            tmp_user_info,
            msg.sender
        );
    }

    function GetResultByGameInfo(uint256 game_id)
        public
        view
        returns (DataDefine.GameResult)
    {
        return GameOperate.GetResultByGameInfo(batch_game[game_id]);
    }

    function GetReward() public CheckCanGetReward {
        // TODO(kingxinwang): add transfer function
        GameOperate.SetGotReward(batch_game[now_game_id], msg.sender);

        require(
            IERC20(GameOperate.GetRequireTokenAddress(batch_game[now_game_id]))
                .transfer(
                    msg.sender,
                    GameOperate.GetUserReward(
                        batch_game[now_game_id],
                        msg.sender
                    )
                ),
            "transfer err"
        );
    }
}
