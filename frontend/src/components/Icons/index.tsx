import { CSSProperties, SVGProps } from "react";

interface LogoProps extends SVGProps<SVGElement> {
  width?: string,
  height?: string,
  fill?: string,
  visibility?: string,
  style?: CSSProperties | undefined;
  flipped?: boolean;
}


export default LogoProps
