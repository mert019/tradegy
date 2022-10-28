import ApiService from "./ApiService";


class LeaderboardService extends ApiService {

    async getLeaderboardData() {

        var myHeaders = new Headers();

        var requestOptions = {
            method: 'GET',
            headers: myHeaders,
            redirect: 'follow'
        };

        const url = this.getBaseApiUrl() + "/api/v1/leaderboard/list";

        return await this.sendApiRequest(url, requestOptions);
    }
}


export default LeaderboardService;
