import React, { useRef, useEffect, useCallback } from "react";

export const Canvas = ({ base64Image, drawingCoords, height, width }) => {
  const canvasRef = useRef(null);

  const draw = useCallback(
    (ctx) => {
      const image = new Image();
      image.onload = () => {
        ctx.drawImage(image, 0, 0);
      };
      image.src = base64Image;
    },
    [base64Image]
  );

  useEffect(() => {
    const canvas = canvasRef.current;
    const context = canvas.getContext("2d");
    context.canvas.width = width;
    context.canvas.height = height;
    context.clearRect(0, 0, width, height);

    // initial draw
    draw(context);
  }, [draw, width, height]);

  const drawFeatures = useCallback((ctx, people) => {
    console.log("features :>> ", people);

    people.forEach((face) => {
      face.forEach((feature) => {
        ctx.beginPath();

        const [a, b, c, d, type] = feature;

        // face
        if (type === 0) {
          const x = b - c / 2;
          const y = a - c / 2;

          ctx.lineWidth = "2";
          ctx.strokeStyle = "red";
          ctx.strokeRect(x, y, c, c);
          ctx.stroke();
        }

        // Eyes
        if (type === 1) {
          const x = b - c / 2;
          const y = a - c / 2;

          ctx.lineWidth = "2";
          ctx.strokeStyle = "blue";
          ctx.strokeRect(x, y, c, c);
          ctx.stroke();
        }

        // Landmarks
        if (type === 2) {
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

    drawFeatures(context, drawingCoords);
  }, [drawingCoords, drawFeatures]);

  return <canvas ref={canvasRef} />;
};
