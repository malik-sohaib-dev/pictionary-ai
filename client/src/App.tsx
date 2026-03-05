import './App.css'
import { useState, useRef, useEffect, useMemo, useCallback } from 'react';

const widthMultiplier = 1.33333333333333;
const initCanvasHeight = 450;
const initCanvasWidth = initCanvasHeight * widthMultiplier;
let drawingHistory: { offSetXTransformed: number, offSetYTransformed: number, color: string, lineWidthTransformed: number }[] = []

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

  const getCanvasDimensions = () => {
    const canvas = canvasRef.current

    let w = initCanvasWidth
    let h = initCanvasHeight
    if (!canvas) return { w, h }

    w = canvas.offsetWidth || initCanvasWidth
    h = canvas.offsetHeight || initCanvasHeight

    return { w, h }
  }


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

      contextInit(canvas)

      recreateHistory()

    } else if (canvas.height / 2 != initCanvasHeight) {
      canvas.width = width * 2;
      canvas.height = height * 2;
      canvas.style.width = `${width}px`;
      canvas.style.height = `${height}px`;

      contextInit(canvas)

      recreateHistory()
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

    contextInit(canvas)
  }

  const contextInit = (canvas: any) => {
    const context = canvas.getContext("2d");
    context.scale(2, 2);
    context.lineCap = "round";
    context.lineJoin = "bevel"
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
    // console.log("Down", { offsetX, offsetY })
    updateDrawingHistory(offsetX, offsetY)
    contextRef.current.beginPath();
    contextRef.current.moveTo(offsetX, offsetY);
    contextRef.current.lineTo(offsetX, offsetY);
    contextRef.current.stroke();
    setIsDrawing(true);
  };

  const draw = ({ nativeEvent }) => {
    if (!isDrawing) return;
    const { offsetX, offsetY } = nativeEvent;
    updateDrawingHistory(offsetX, offsetY)
    // console.log({ offsetX, offsetY })
    contextRef.current.lineTo(offsetX, offsetY);
    contextRef.current.stroke();
  };

  const showHistory = () => {
    console.log(drawingHistory)
  }

  const updateDrawingHistory = (offSetX: number, offSetY: number) => {

    let { w, h } = getCanvasDimensions();

    if (offSetX === -1) {
      drawingHistory.push({ offSetXTransformed: offSetX, offSetYTransformed: offSetY, color, lineWidthTransformed: lineWidth })
    } else {
      drawingHistory.push({ offSetXTransformed: offSetX / w, offSetYTransformed: offSetY / h, color, lineWidthTransformed: lineWidth / w })
    }
  }

  const resetHistory = () => {
    drawingHistory = []
  }

  const recreateHistory = () => {
    if (!drawingHistory[0]) return

    const { w, h } = getCanvasDimensions();

    contextRef.current.beginPath();
    contextRef.current.moveTo(drawingHistory[0].offSetXTransformed * w, drawingHistory[0].offSetYTransformed * h);


    drawingHistory.map((history, ind) => {
      if (history.offSetXTransformed === -1 && drawingHistory[ind + 1]?.offSetXTransformed) {
        contextRef.current.beginPath();
        contextRef.current.moveTo(drawingHistory[ind + 1].offSetXTransformed * w, drawingHistory[ind + 1].offSetYTransformed * h);

      } else if (!(history.offSetXTransformed === -1)) {
        contextRef.current.lineWidth = history.lineWidthTransformed * w;
        contextRef.current.strokeStyle = history.color;
        contextRef.current.lineTo(history.offSetXTransformed * w, history.offSetYTransformed * h);
        contextRef.current.stroke();
      }
    })
  }

  const finishDrawing = () => {
    updateDrawingHistory(-1, -1)
    setIsDrawing(false);
  };

  const changeColor = (value: string) => {
    if (value === color) return;
    setColor(value);
  }

  const changeLineWidth = (value: string | number) => {
    const valueInt = Number(value)
    if (valueInt === lineWidth || valueInt < 1 || valueInt > 50) return;
    setLineWidth(valueInt);
  }

  const clearAll = () => {
    canvasRef.current.getContext("2d").reset()
    canvasInit()
  }

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
            return (<div key={val} className={`w-7 h-7 mr-1.5 cursor-pointer ${color === val ? 'border-4' : ''}`} style={{ backgroundColor: val }} onClick={() => changeColor(val)} />)
          })
        }
      </div>

      {/* Line width selection */}
      <div>
        <input type='range' min={1} max={50} className='w-[600px]' value={lineWidth} onChange={(e) => changeLineWidth(e.target.value)} /> {lineWidth}
      </div>

      <button className='border bg-amber-200 p-2 cursor-pointer shadow-xl shadow-black' onClick={clearAll}>Clear All</button>

      <button className='border bg-amber-200 p-2 cursor-pointer shadow-xl shadow-black' onClick={showHistory}>Show History</button>

      <button className='border bg-amber-200 p-2 cursor-pointer shadow-xl shadow-black' onClick={recreateHistory}>Recreate History</button>

      <button className='border bg-amber-200 p-2 cursor-pointer shadow-xl shadow-black' onClick={resetHistory}>Reset History</button>

    </div>
  )
}

export default App
