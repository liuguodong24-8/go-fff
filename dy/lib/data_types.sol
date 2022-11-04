// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

library DataDefine {
    enum GameStatus {
        kNoStart,
        kIng,
        kEnd
    }

    enum GameResult {
        kWin,
        kDraw,
        kFail
    }

    enum PositionDefine {
        kFirst,
        kSecond
    }

    struct UserInfo {
        uint256 amount;
        GameResult choose_result;
        bool got_reward;
    }

    struct GameInfo {
        address token_address;
        string first_team;
        string second_team;
        uint256 start_time;
        uint256 end_time;
        // init zero; end_time can change
        uint16 first_team_scores;
        uint16 second_team_scores;
        // game status
        GameStatus game_status;
        // total amount
        uint256 total_stake_amount;
        uint256 win_stake_amount;
        uint256 fail_stake_amount;
        uint256 draw_stake_amount;
    }

    // need get some gee;
    struct DistrubuteMode {
        uint256 fee_rate;
        uint256 fee_base;
    }

    // one game struct info
    struct OneGameInfo {
        DataDefine.GameInfo game_info;
        DataDefine.DistrubuteMode fee_info;
        mapping(address => DataDefine.UserInfo) user_infos;
    }
}
