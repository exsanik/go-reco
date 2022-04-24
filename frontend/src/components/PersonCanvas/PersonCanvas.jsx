import React, { useRef, useEffect, useCallback } from "react";
import { TYPE_EYE, TYPE_FACE, TYPE_FEATURE } from "src/config";

export const PersonCanvas = ({ drawingCoords, height, width, className }) => {
  const canvasRef = useRef(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    const context = canvas.getContext("2d");
    context.canvas.width = width;
    context.canvas.height = height;
  }, [width, height]);

  const drawFeatures = useCallback((ctx, people) => {
    people.forEach((face) => {
      face.forEach((feature) => {
        ctx.beginPath();

        const [a, b, length, type] = feature;

        if (type === TYPE_FACE) {
          const x = b - length / 2;
          const y = a - length / 2;

          ctx.lineWidth = "2";
          ctx.strokeStyle = "red";
          ctx.strokeRect(x, y, length, length);
          ctx.stroke();
        }

        if (type === TYPE_EYE) {
          const x = b - length / 2;
          const y = a - length / 2;

          ctx.lineWidth = "2";
          ctx.strokeStyle = "blue";
          ctx.strokeRect(x, y, length, length);
          ctx.stroke();
        }

        if (type === TYPE_FEATURE) {
          ctx.arc(b, a, 4, 0, 2 * Math.PI, false);
          ctx.lineWidth = "2";
          ctx.strokeStyle = "green";
          ctx.stroke();
        }

        ctx.closePath();
      });
    });
  }, []);

  useEffect(() => {
    const canvas = canvasRef.current;
    const context = canvas.getContext("2d");

    context.clearRect(0, 0, canvas.width, canvas.height);
    drawFeatures(context, drawingCoords);
  }, [drawingCoords, drawFeatures]);

  return <canvas ref={canvasRef} className={className} />;
};
