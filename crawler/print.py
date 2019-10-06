import torch
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

edges = pd.read_csv("teste.csv")


test = torch.load("ptbr.pth")
nice = list(zip(*[test["objects"], test["embeddings"]]))
di = dict(nice)
teste = list(zip(*nice))
X = list(map(lambda x: x[0].numpy(), teste[1]))
Y = list(map(lambda x: x[1].numpy(), teste[1]))
plt.figure(figsize=(100, 100))
plt.scatter(
    X, Y, s=80*(1-np.array(list(map(lambda x: x.norm(p=2), teste[1])))))
for i, label in enumerate(teste[0]):
    if edges.loc[edges["id1"] == label]["id2"].any() == "nome" or label == "nome":
        plt.annotate(label, (X[i], Y[i]))
       # if di[label].norm(2) < 0.8:
       #         plt.annotate(label, (X[i], Y[i]))

plt.show()
