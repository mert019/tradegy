import { useState, useEffect } from 'react';

import LeaderboardService from '../services/LeaderboardService';


const LeaderboardPage = () => {

    const [leaderboardData, setLeaderboardData] = useState(null);

    const leaderboardService = new LeaderboardService();


    const getLeaderboardData = async () => {

        let response = await leaderboardService.getLeaderboardData();

        // Error
        if (response === null) {
            alert.error("Ooops something went wrong.")
        }
        // Success
        else if (200 <= response.Status && response.Status <= 299) {
            setLeaderboardData(response.Payload);
        }
        // Error
        else {
            if (response.Message.length > 0) {
                alert.error(response.Message)
            }
        }
    }


    useEffect(() => {
        getLeaderboardData();
    }, [])


    return (
        <div className="container mt-3">
            <div className="d-flex justify-content-between flex-wrap">
                <span className="h4">Leaderboard</span>
                <span>
                    Updated At:
                    {leaderboardData !== null && (new Date(leaderboardData.updated_at)).toLocaleString("en-US")}
                </span>
            </div>

           
            <table class="table">
                <thead>
                    <tr>
                        <th scope="col" className="text-left">#</th>
                        <th scope="col" className="text-right">Username</th>
                        <th scope="col" className="text-right">Total Wealth (USD)</th>
                    </tr>
                </thead>
                <tbody>
                    {leaderboardData !== null &&
                        leaderboardData.leaderboard_items.map((val, index) => {
                            return <tr key={"leaderboardData_" + index}>
                                <th scope="row" className="text-left">{index + 1}</th>
                                <td className="text-right">{val.user_name}</td>
                                <td className="text-right">{val.total_usd_amount.toFixed(2)} $</td>
                            </tr>
                        })
                    }
                </tbody>
            </table>
        </div>
    )
}


export default LeaderboardPage;
