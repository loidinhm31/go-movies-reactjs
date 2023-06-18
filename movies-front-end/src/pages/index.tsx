import { getDefaultLayout } from "src/components/Layout/Layout";
import Home from "src/pages/home";

function App() {
    return <Home />;
}

App.getLayout = getDefaultLayout;

export default App;
