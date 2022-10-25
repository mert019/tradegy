const Footer = () => {
    return (
        <footer className="p-3 bg-dark text-white text-center">
            <div>&#169; 2022 Copyright: {process.env.REACT_APP_COMPANY_NAME}</div>
        </footer>
    )
}

export default Footer;
