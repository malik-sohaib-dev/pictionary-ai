import './App.css'
import { useState, useRef, useEffect, useMemo, useCallback } from 'react';

const widthMultiplier = 1.33333333333333;
const initCanvasHeight = 450;
const initCanvasWidth = initCanvasHeight * widthMultiplier;

function App() {
  const canvasRef = useRef(null);
  const contextRef = useRef(null);
  const [isDrawing, setIsDrawing] = useState(false);
  const [color, setColor] = useState('black');
  const [lineWidth, setLineWidth] = useState(10);
  const colors = useMemo(() => {
    return [
      "black",
      "red",
      "yellow",
      "pink",
      "blue",
      "#AAAAAA",
      "white"
    ]
  }, []);

  useEffect(() => {
    canvasInit()
  }, [])


  useEffect(() => {

    window.addEventListener("resize", windowResize)

    return () => {
      window.removeEventListener("resize", windowResize)
    }
  });


  const windowResize = useCallback(() => {
    const canvas = canvasRef.current;
    let width = initCanvasWidth;
    let height = initCanvasHeight;

    if (window.innerWidth < initCanvasWidth) {
      width = window.innerWidth
      height = window.innerWidth / widthMultiplier;
      canvas.width = width * 2;
      canvas.height = height * 2;
      canvas.style.width = `${width}px`;
      canvas.style.height = `${height}px`;

    } else if (canvas.height/2 != initCanvasHeight) {
      canvas.width = width * 2;
      canvas.height = height * 2;
      canvas.style.width = `${width}px`;
      canvas.style.height = `${height}px`;
    }
  }, [])

  const canvasInit = () => {
    let width = initCanvasWidth;
    let height = initCanvasHeight;
    if (window.innerWidth < initCanvasWidth) {
      width = window.innerWidth
      height = window.innerWidth / widthMultiplier;
    }
    const canvas = canvasRef.current;
    // Set internal resolution to match display size
    canvas.width = width * 2;
    canvas.height = height * 2;
    canvas.style.width = `${width}px`;
    canvas.style.height = `${height}px`;

    const context = canvas.getContext("2d");
    context.scale(2, 2);
    context.lineCap = "round";
    context.strokeStyle = color;
    context.lineWidth = lineWidth;
    contextRef.current = context;
  }

  useEffect(() => {
    contextRef.current.lineWidth = lineWidth;
    contextRef.current.strokeStyle = color;
  }, [color, lineWidth])

  const startDrawing = ({ nativeEvent }) => {
    const { offsetX, offsetY } = nativeEvent;
    contextRef.current.beginPath();
    contextRef.current.moveTo(offsetX, offsetY);
    contextRef.current.lineTo(offsetX, offsetY);
    contextRef.current.stroke();
    setIsDrawing(true);
  };

  const finishDrawing = () => {
    contextRef.current.closePath();
    setIsDrawing(false);
  };

  const colorChange = (value: string) => {
    if (value === color) return;
    setColor(value);
  }

  const lineWidthChange = (value: string) => {
    const valueInt = Number(value)
    if (valueInt === lineWidth || valueInt < 1 || valueInt > 50) return;
    setLineWidth(valueInt);
  }

  const clearAll = () => {
    canvasRef.current.getContext("2d").reset()
    canvasInit()
  }

  const draw = ({ nativeEvent }) => {
    if (!isDrawing) return;
    const { offsetX, offsetY } = nativeEvent;
    contextRef.current.lineTo(offsetX, offsetY);
    contextRef.current.stroke();
  };
  return (
    <div>
      <canvas
        ref={canvasRef}
        onMouseDown={startDrawing}
        onMouseUp={finishDrawing}
        onMouseMove={draw}
        className="bg-white"
      />

      {/* Color selection */}
      <div className='flex m-4'>
        {
          colors.map((val) => {
            return (<div key={val} className={`w-7 h-7 mr-1.5 cursor-pointer ${color === val ? 'border-4' : ''}`} style={{ backgroundColor: val }} onClick={() => colorChange(val)} />)
          })
        }
      </div>

      {/* Line width selection */}
      <div>
        <input type='range' min={1} max={50} className='w-[600px]' value={lineWidth} onChange={(e) => lineWidthChange(e.target.value)} /> {lineWidth}
      </div>

      <button onClick={clearAll}>Clear All</button>

    </div>
  )
}

export default App
