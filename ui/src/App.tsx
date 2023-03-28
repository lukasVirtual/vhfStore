import SoftwareList from "./components/Software/software.tsx"
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

function App() {
  return (
  <ThemeProvider theme={darkTheme}>
  <CssBaseline />
      <SoftwareList />
  </ThemeProvider>

  );
}

export default App;
