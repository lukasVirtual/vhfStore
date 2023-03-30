import { createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import SoftwareList from "./components/Software/software";
import PrimarySearchAppBar from "./AppBar";

const darkTheme = createTheme({
  palette: {
    mode: "dark",
  },
});

const res = await fetch("http://127.0.0.1:8080/getApps")
const apps: [] = await res.json()
console.log(apps)


function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <div>
        <PrimarySearchAppBar />
        {apps.map((v) => (
          <SoftwareList fn={v.ProductName} fd={v.FileDescription} cn={v.CompanyName} v={v.ProductVersion} />
        ))}
      </div>
    </ThemeProvider>
  );
}

export default App;
