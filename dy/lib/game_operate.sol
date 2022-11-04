// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./data_types.sol";
import "./const.sol";

library GameOperate {
    function InitGameInfo(
        DataDefine.GameInfo storage now_game_info,
        string[] memory team_infos,
        uint256[] memory time_infos,
        address stake_address
    ) external {
        now_game_info.first_team = team_infos[Const.kFirst];
        now_game_info.second_team = team_infos[Const.kSecond];

        now_game_info.start_time = time_infos[Const.kFirst];
        now_game_info.end_time = time_infos[Const.kSecond];

        // status;
        now_game_info.game_status = DataDefine.GameStatus.kNoStart;

        // token address;
        now_game_info.token_address = stake_address;
    }

    function InitFeeInfos(
        DataDefine.DistrubuteMode storage now_fee_info,
        uint256[] memory fee_infos
    ) external {
        now_fee_info.fee_rate = fee_infos[Const.kFirst];
        now_fee_info.fee_base = fee_infos[Const.kSecond];
    }

    function StartGame(DataDefine.GameInfo storage game_info) external {
        game_info.game_status = DataDefine.GameStatus.kIng;
    }

    function StopGame(DataDefine.GameInfo storage game_info) external {
        game_info.game_status = DataDefine.GameStatus.kEnd;
    }

    function GetResultByGameInfo(DataDefine.OneGameInfo storage one_game)
        external
        view
        returns (DataDefine.GameResult)
    {
        return GameOperate._GetResultByGameInfo(one_game);
    }

    function _GetResultByGameInfo(DataDefine.OneGameInfo storage one_game)
        internal
        view
        returns (DataDefine.GameResult)
    {
        if (
            one_game.game_info.first_team_scores ==
            one_game.game_info.second_team_scores
        ) {
            return DataDefine.GameResult.kDraw;
        }

        if (
            one_game.game_info.first_team_scores >
            one_game.game_info.second_team_scores
        ) {
            return DataDefine.GameResult.kWin;
        }

        if (
            one_game.game_info.first_team_scores <
            one_game.game_info.second_team_scores
        ) {
            return DataDefine.GameResult.kFail;
        }

        return DataDefine.GameResult.kFail;
    }

    function GetGameUserInfo(
        DataDefine.OneGameInfo storage one_game,
        address user_address
    ) public view returns (DataDefine.UserInfo memory) {
        return one_game.user_infos[user_address];
    }

    function SetGotReward(
        DataDefine.OneGameInfo storage one_game,
        address user_address
    ) external {
        one_game.user_infos[user_address].got_reward = true;
    }

    function SetGameScores(
        uint16[2] memory scores,
        DataDefine.GameInfo storage game_info
    ) external {
        game_info.first_team_scores = scores[Const.kFirst];
        game_info.second_team_scores = scores[Const.kSecond];
    }

    function GetGameStatus(DataDefine.GameInfo memory game_info)
        external
        pure
        returns (DataDefine.GameStatus)
    {
        return game_info.game_status;
    }

    function GetRequireTokenAddress(DataDefine.OneGameInfo storage one_game)
        external
        view
        returns (address)
    {
        return one_game.game_info.token_address;
    }

    function AddGameStakeUser(
        DataDefine.OneGameInfo storage one_game,
        DataDefine.UserInfo memory user_info,
        address user_address
    ) external {
        // add user info to one_game struct
        one_game.user_infos[user_address] = user_info;

        // gameinfo add total amount
        one_game.game_info.total_stake_amount += user_info.amount;

        // judge win || fail || draw
        if (user_info.choose_result == DataDefine.GameResult.kWin) {
            one_game.game_info.win_stake_amount += user_info.amount;
        }

        if (user_info.choose_result == DataDefine.GameResult.kDraw) {
            one_game.game_info.draw_stake_amount += user_info.amount;
        }

        if (user_info.choose_result == DataDefine.GameResult.kFail) {
            one_game.game_info.fail_stake_amount += user_info.amount;
        }
    }

    function GetUserReward(
        DataDefine.OneGameInfo storage one_game,
        address user_address
    ) external view returns (uint256) {
        uint256 user_result_map_amount = 0;

        if (_GetResultByGameInfo(one_game) == DataDefine.GameResult.kDraw) {
            user_result_map_amount = one_game.game_info.draw_stake_amount;
        }

        if (_GetResultByGameInfo(one_game) == DataDefine.GameResult.kWin) {
            user_result_map_amount = one_game.game_info.win_stake_amount;
        }

        if (_GetResultByGameInfo(one_game) == DataDefine.GameResult.kFail) {
            user_result_map_amount = one_game.game_info.fail_stake_amount;
        }

        // TODO(kingxinwang): safe math
        return
            (one_game.user_infos[user_address].amount *
                one_game.game_info.total_stake_amount) / user_result_map_amount;
    }
}
