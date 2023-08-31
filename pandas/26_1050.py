import pandas as pd

def actors_and_directors(actor_director: pd.DataFrame) -> pd.DataFrame:
    df = actor_director.groupby(["actor_id", "director_id"]).size().reset_index(name="count")

    return df[df["count"] >= 3][["actor_id", "director_id"]]


if __name__ == "__main__":
    data = [{"actor_id":"","director_id":"","timestamp": ""},
            {"actor_id": 1, "director_id": 1, "timestamp": 0},
            {"actor_id": 1, "director_id": 1, "timestamp": 1},
            {"actor_id": 1, "director_id": 1, "timestamp": 2},
            {"actor_id": 1, "director_id": 2, "timestamp": 3},
            {"actor_id": 1, "director_id": 2, "timestamp": 4},
            {"actor_id": 2, "director_id": 1, "timestamp": 5},
            {"actor_id": 2, "director_id": 1, "timestamp": 0},
            ]
    df = pd.DataFrame(data)
    df = actors_and_directors(df)
    print(df)
