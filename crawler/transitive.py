import pandas as pd
hypers = {}
edges = set()
def import_data() -> None:
    return pd.read_csv("appended.csv")
def get_hypers(word: str, hypers: dict, m) -> bool:
    if word not in hypers:
        return 
    for hyper in hypers[word]:
        edges.add((word, hyper))
        if hyper in hypers.keys():
            for hy in hypers[hyper]:
                edges.add((word, hy))



if __name__ == "__main__":
    data = import_data()
    
    for _, line in data.iterrows():
        if line['id1'] in hypers.keys():
            hypers[line["id1"]] += [line["id2"]]
        else:
            hypers[line["id1"]] = [line["id2"]]
    for word in hypers.keys():
        get_hypers(word, hypers, 0)
    df = pd.DataFrame(list(edges), columns=["id1", "id2"])
    df["weight"] = 1
    print(df)
    df.to_csv("teste.csv", index=False)
