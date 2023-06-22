import { getDefaultLayout } from "@/components/Layout/Layout";
import Home from "@/pages/home";

function App() {
  return <Home />;
}

App.getLayout = getDefaultLayout;

export default App;
