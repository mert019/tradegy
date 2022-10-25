import PieChart, { Series, Label, Connector, Size, Legend, } from 'devextreme-react/pie-chart';



const WealthInformation = ({ wealthInfo }) => {

    
    const renderCenter = (pieChart) => {
        const totalUsdAmount = Object.keys(wealthInfo).length === 0 ? 0 : wealthInfo.reduce((accumulator, object) => {
            return accumulator + object.usd_amount;
        }, 0);
        return (
            <svg>
                <circle cx="100" cy="100" r={pieChart.getInnerRadius() - 6} fill="#fff"></circle>
                <text textAnchor="middle" x="100" y="110" style={{ fontSize: 18, fill: '#333' }}>
                    <tspan style={{ fontWeight: 600 }}>
                        {"$ " + totalUsdAmount.toLocaleString("en-US")}
                    </tspan>
                </text>
            </svg>
        );
    }


    return (
        <PieChart
            id="pie"
            type="doughnut"
            dataSource={wealthInfo}
            palette="Bright"
            centerRender={renderCenter}
        >
            <Legend
                orientation="horizontal"
                itemTextPosition="right"
                horizontalAlignment="center"
                verticalAlignment="bottom"
                columnCount={4} />
            <Series
                argumentField="name"
                valueField="usd_amount"
            >
                <Label visible={true} customizeText={(args) => { return `${parseFloat(args.percentText).toLocaleString("en-US")}%` }}>
                    <Connector visible={true} width={1} />
                </Label>
            </Series>
            <Size width="100%" />
        </PieChart>
    )
}

export default WealthInformation
