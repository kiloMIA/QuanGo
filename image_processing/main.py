import cv2
import numpy as np
from imutils.perspective import four_point_transform
from matplotlib import pyplot as plt

def main():
    image = cv2.imread("dataset/dataset/6/Zrzut ekranu 2020-04-11 o 15.43.19.png")

    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    blurred = cv2.GaussianBlur(gray, (5, 5), 0)
    edges = cv2.Canny(blurred, 75, 200)

    cv2.imwrite('transformed_image.png', edges)

if __name__ == "__main__":
    main()

