import { Button, Card } from "@mui/material"

function SoftwareList<T>(props: T) {
  return (
    <Card className="flex justify-around items-center m-10 rounded-xl h-28">
      <h1 className="text-xl text-left">{props.fn}</h1 >
      <h1 className="text-l">{props.fd}</h1 >
      <h1 className="text-xl text-justify"> {props.cn}</h1 >
      <h1 className="text-xl text-right"> {props.v}</h1 >
      <Button variant="contained" >Install</Button>
    </Card >
  )
}

export default SoftwareList;
