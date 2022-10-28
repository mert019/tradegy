const MainPage = () => {
  return (
    <>
      <div className="mb-4 mt-4">
        <h2 className="text-center">{process.env.REACT_APP_COMPANY_NAME}: Mock Trading Platform</h2>
        <p className="text-center">If you don't have a strategy, trading becomes tragedy.</p>
      </div>

      <div className="container">
        <div className="mb-3">
          <coingecko-coin-ticker-widget coin-id="bitcoin" currency="usd" locale="en"></coingecko-coin-ticker-widget>
        </div>
        <div className="mb-3">
          <coingecko-coin-ticker-widget coin-id="dogecoin" currency="usd" locale="en"></coingecko-coin-ticker-widget>
        </div>
      </div>
    </>
  )
}

export default MainPage;
