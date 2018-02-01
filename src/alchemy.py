from sqlalchemy import create_engine, MetaData, Table, Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base


DeclaredBaseModel = declarative_base()


class YakSokModel(DeclaredBaseModel):
    __tablename__ = "yaksoks"
    id = Column(Integer, primary_key=True)
    cmd = Column(String(256), nullable=False)
    path = Column(String(256), nullable=False)
    uid = Column(Integer, nullable=False, index=True)
    at = Column(String(10), nullable=False)
    interval = Column(Integer, nullable=True)
    routine = Column(String(10), nullable=True)
    tags = Column(String(256), nullable=True)
    name = Column(String(128), nullable=True, index=True)

    def __init__(self, cmd: str, path: str, uid: int, at: str, routine="", interval=0, tags="", name=""):
        # self.id = id
        self.cmd = cmd
        self.path = path
        self.uid = uid
        self.at = at
        self.routine = routine
        self.interval = interval
        self.tags = tags
        self.name = name

    pass


class YakSokController:
    metadata = None

    @staticmethod
    def factory():
        engine = create_engine('sqlite:///yaksok.db', echo=True)
        YakSokController.metadata = MetaData(bind=engine)
        YakSokModel.__table__.create(bind=engine)
        ysctrl = YakSokController()
        return ysctrl

    def __init__(self):
        self.table = {
            'yaksoks': Table(YakSokModel.__tablename__, self.metadata, autoload=True)
        }
        pass

    pass