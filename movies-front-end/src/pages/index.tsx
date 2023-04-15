import {getDefaultLayout} from "../components/Layout";
import Home from "./home";


function App() {
    return (
        <Home />
    );
}

App.getLayout = getDefaultLayout;

export default App;
